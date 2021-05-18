package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

type process struct {
	startBalances  []uint64
	finishBalances []uint64
	startedAt      time.Time
	mustStopAfter  time.Time
	nextDelayAfter time.Time
	txLimit        uint64
	sentTx         uint64
	valuesCount    uint64
	finish         bool
	interactive    bool
	timeSpent      time.Duration
	delayUpTo      time.Duration
	delayNow       time.Duration
	s              sync.Mutex
}

func init() {
	rand.Seed(time.Now().Unix())
}

func sendTxSpender(wg *sync.WaitGroup, from *wallet, workset []*wallet, amount uint64, nodeUrl string, chainId string, p *process) {
	defer wg.Done()
	var to *wallet
	var tx string
	var err error
	//log.Println("All Txs for wallet", from.address, " will be sent to node:", nodeUrl)
	retryInt := 10 * time.Millisecond
	l := len(workset)
	for {
		// select target wallet
		if l == 2 {
			// nothing to rand
			if workset[0].address == from.address {
				to = workset[1]
			} else {
				to = workset[0]
			}
		} else {
			to = workset[rand.Intn(l)]
			if to.address == from.address {
				continue // try again
			}
		}

		// check source balance
		from.s.Lock()
		if from.balance < amount+feesAmount {
			log.Println("wallet", from.address, " out of tokens:", from.balance)
			from.s.Unlock()
			return // source wallet out of tokens
		}
		tx = getSignedSendTx(from.address, to.address, amount, "", from.privKey, chainId, from.accountNumber, from.sequence)
		from.s.Unlock()
		// check again right before broadcast to prevent excess Txs
		p.s.Lock()
		if p.finish {
			p.s.Unlock()
			return
		}
		p.s.Unlock()

		_, err = broadcastTx(tx, nodeUrl, "sync")
		seqRetries := 10
		for err == ErrMempoolIsFull || err == ErrTooManyOpenFiles || (err == ErrSequenceWrong && seqRetries > 0) {
			// try to fix sequence
			if err == ErrSequenceWrong {
				from.s.Lock()
				from.sequence--
				tx = getSignedSendTx(from.address, to.address, amount, "", from.privKey, chainId, from.accountNumber, from.sequence)
				from.s.Unlock()
				seqRetries--
				time.Sleep(2 * time.Millisecond)
			} else {
				// wait & retry
				time.Sleep(retryInt)
				if retryInt < 100*time.Millisecond {
					retryInt *= 2 // progressive pause, but not longer 100ms
				}
			}
			_, err = broadcastTx(tx, nodeUrl, "sync")
		}
		retryInt = 10 * time.Millisecond // reset progressive pause
		if err != nil {
			log.Println("broadcast FAIL for", from.address, ", with sequence:", from.sequence, ", err:", err)
			time.Sleep(time.Second)
			return
		}
		// calc balances
		from.s.Lock()
		from.balance -= amount + feesAmount
		from.sequence++
		from.s.Unlock()
		to.s.Lock()
		to.balance += amount
		to.s.Unlock()
		// check process state
		if p.CalcTx() {
			return
		}
		// current delay
		if p.delayUpTo > 0 {
			time.Sleep(p.delayNow)
		}
		// default minimal delay to prevent flooding of mempool with too fast requests
		time.Sleep(2 * time.Millisecond)
	}
}

func rapidIntakeSpender(wg *sync.WaitGroup, from *wallet, sc common.Address, mode, ic uint64, nodeUrl string, chainId string, p *process) {
	defer wg.Done()
	var tx string
	var err error
	arraySlot := common.HexToHash("0x0")
	singleSlot := common.HexToHash("0x1")
	retryInt := 10 * time.Millisecond
	nextCall := mode
	if mode == 6 {
		nextCall = 4
	}
	for {
		value := common.BigToHash(big.NewInt(rand.Int63()))
		values := make([]common.Hash, ic)
		for i := range values {
			values[i] = common.BigToHash(big.NewInt(rand.Int63()))
		}
		idx := common.BigToHash(big.NewInt(rand.Int63n(32768))) // total array length up to 32K
		// check source balance
		from.s.Lock()
		if from.balance < feesAmount {
			log.Println("wallet", from.address, " out of tokens:", from.balance)
			from.s.Unlock()
			return // source wallet out of tokens
		}
		switch nextCall {
		case 4:
			tx = getSignedSetValueTx(from.address, sc, singleSlot, value, "", from.privKey, chainId, from.accountNumber, from.sequence)
		case 5:
			tx = getSignedSetArrayValuesTx(from.address, sc, arraySlot, idx, values, "", from.privKey, chainId, from.accountNumber, from.sequence)
		}
		from.s.Unlock()
		// check again right before broadcast to prevent excess Txs
		p.s.Lock()
		if p.finish {
			p.s.Unlock()
			return
		}
		p.s.Unlock()

		_, err = broadcastTx(tx, nodeUrl, "sync")
		seqRetries := 10
		for err == ErrMempoolIsFull || err == ErrTooManyOpenFiles || (err == ErrSequenceWrong && seqRetries > 0) {
			// try to fix sequence
			if err == ErrSequenceWrong {
				from.s.Lock()
				from.sequence--
				switch nextCall {
				case 4:
					tx = getSignedSetValueTx(from.address, sc, singleSlot, value, "", from.privKey, chainId, from.accountNumber, from.sequence)
				case 5:
					tx = getSignedSetArrayValuesTx(from.address, sc, arraySlot, idx, values, "", from.privKey, chainId, from.accountNumber, from.sequence)
				}
				from.s.Unlock()
				seqRetries--
				time.Sleep(2 * time.Millisecond)
			} else {
				// wait & retry
				time.Sleep(retryInt)
				if retryInt < 100*time.Millisecond {
					retryInt *= 2 // progressive pause, but not longer 100ms
				}
			}
			_, err = broadcastTx(tx, nodeUrl, "sync")
		}
		retryInt = 10 * time.Millisecond // reset progressive pause
		if err != nil {
			log.Println("broadcast FAIL for", from.address, ", with sequence:", from.sequence, ", err:", err)
			time.Sleep(time.Second)
			return
		}
		// calc balances
		from.s.Lock()
		from.balance -= feesAmount
		from.sequence++
		from.s.Unlock()
		// check process state
		if p.CalcTx() {
			return
		}
		// current delay
		if p.delayUpTo > 0 {
			time.Sleep(p.delayNow)
		}
		// default minimal delay to prevent flooding of mempool with too fast requests
		time.Sleep(2 * time.Millisecond)
		// next call will be..
		if mode == 6 {
			// flip-flop calls
			if nextCall == 4 {
				nextCall = 5
			} else {
				nextCall = 4
			}
		} else {
			nextCall = mode
		}
	}
}

func (p *process) CalcTx() bool {
	p.s.Lock()
	defer p.s.Unlock()
	p.sentTx++
	if p.finish {
		return true
	}
	// check Txs limit
	if p.txLimit > 0 && p.sentTx >= p.txLimit {
		p.finish = true
	}
	// check time limit
	if !p.mustStopAfter.IsZero() && time.Now().After(p.mustStopAfter) {
		p.finish = true
	}
	// final stat
	if p.finish {
		if p.interactive {
			fmt.Print("\r\x1b[2K") // clear line
		}
		timeSpent := time.Now().Sub(p.startedAt)
		rps := p.sentTx
		if timeSpent.Seconds() > 1 {
			rps /= uint64(timeSpent.Seconds())
		}
		if p.interactive {
			fmt.Print("\r")
		} else {
			fmt.Print("\n")
		}
		fmt.Print("DONE - ", p.sentTx, " Txs was sent, ", timeSpent, " time was spent, rps - ", rps, "\n")
		return true
	}
	// change delay each 1000 txs
	if p.delayUpTo > 0 && time.Now().After(p.nextDelayAfter) {
		p.delayNow = time.Duration(rand.Int63n(int64(p.delayUpTo)))
		p.nextDelayAfter = time.Now().Add(10 * p.delayNow)
	}
	// update progress line each 100 Txs
	if p.sentTx%100 == 0 {
		if p.interactive {
			fmt.Print("\r\x1b[2K") // clear line
		}
		if p.interactive {
			fmt.Print("\r")
		} else {
			fmt.Print("\n")
		}
		fmt.Print("Progress - ", p.sentTx, " Txs was sent,")
		if p.txLimit > 0 {
			fmt.Print(" another ", p.txLimit-p.sentTx, " Txs")
		}
		if !p.mustStopAfter.IsZero() {
			if p.txLimit > 0 {
				fmt.Print(" or")
			}
			fmt.Print(" ", p.mustStopAfter.Sub(time.Now()))
		}
		fmt.Print(" left")
		if p.delayUpTo > 0 {
			fmt.Printf(", current delay: %v", p.delayNow)
		}
		if p.interactive {
			fmt.Print("\r")
		}
	}
	return p.finish
}
