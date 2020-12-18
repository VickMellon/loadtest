package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type process struct {
	startBalances  []uint64
	finishBalances []uint64
	startedAt      time.Time
	mustStopAfter  time.Time
	txLimit        uint64
	sentTx         uint64
	finish         bool
	timeSpent      time.Duration
	s              sync.Mutex
}

func init() {
	rand.Seed(time.Now().Unix())
}

func spender(wg *sync.WaitGroup, from *wallet, workset []*wallet, amount uint64, nodeUrl string, chainId string, p *process) {
	defer wg.Done()
	var to *wallet
	var tx string
	var err error
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
		for err == ErrMempoolIsFull || err == ErrTooManyOpenFiles {
			// wait & retry
			time.Sleep(retryInt)
			if retryInt < 100*time.Millisecond {
				retryInt *= 2 // progressive pause, but not longer 100ms
			}
			_, err = broadcastTx(tx, nodeUrl, "sync")
		}
		retryInt = 10 * time.Millisecond // reset progressive pause
		if err != nil {
			log.Println("broadcast FAIL, with sequence:", from.sequence, " tx:", tx)
			break
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
		time.Sleep(2 * time.Millisecond) // prevent flooding of mempool with too fast requests
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
		fmt.Print("\r\x1b[2K") // clear line
		timeSpent := time.Now().Sub(p.startedAt)
		rps := p.sentTx
		if timeSpent.Seconds() > 1 {
			rps /= uint64(timeSpent.Seconds())
		}
		fmt.Print("\rDONE - ", p.sentTx, " Txs was sent, ", timeSpent, " time was spent, rps - ", rps, "\n")
		return true
	}
	// update progress line each 100 Txs
	if p.sentTx%100 == 0 {
		fmt.Print("\r\x1b[2K") // clear line
		fmt.Print("\rProgress - ", p.sentTx, " Txs was sent,")
		if p.txLimit > 0 {
			fmt.Print(" another ", p.txLimit-p.sentTx, " Txs")
		}
		if !p.mustStopAfter.IsZero() {
			if p.txLimit > 0 {
				fmt.Print(" or")
			}
			fmt.Print(" ", p.mustStopAfter.Sub(time.Now()))
		}
		fmt.Print(" left\r")
	}
	return p.finish
}
