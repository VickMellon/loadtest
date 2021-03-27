package main

import (
	"github.com/VickMellon/loadtest/MiniStore"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"math/rand"
	"sync"
	"time"
)

const (
	gasPrice  = uint64(1)
	gasWanted = uint64(50000)
)

func sc_caller(wg *sync.WaitGroup, from *wallet, instance *MiniStore.MiniStore, p *process, mode uint64) {
	var err error
	defer wg.Done()
	auth := from.auth
	nextCall := mode
	if mode == 3 {
		nextCall = 1
	}
	for {
		// check source balance
		from.s.Lock()
		if from.balance <= gasWanted {
			log.Println("wallet", from.address, " out of tokens:", from.balance)
			from.s.Unlock()
			return // source wallet out of tokens
		}
		auth.GasLimit = gasWanted
		auth.Nonce = big.NewInt(int64(from.sequence))
		from.s.Unlock()
		// check again right before broadcast to prevent excess Txs
		p.s.Lock()
		if p.finish {
			p.s.Unlock()
			return
		}
		p.s.Unlock()
		// call SC method
		amount := big.NewInt(rand.Int63())
		if nextCall == 1 {
			_, err = instance.SetNumberValue(auth, amount)
			if err != nil {
				log.Println("call SetNumberValue() FAIL, with sequence:", from.sequence, " err:", err)
				time.Sleep(time.Second)
				continue
			}
		} else if nextCall == 2 {
			_, err = instance.AddValue(auth, amount)
			if err != nil {
				log.Println("call AddValue() FAIL, with sequence:", from.sequence, " err:", err)
				time.Sleep(time.Second)
				continue
			}
			p.s.Lock()
			p.valuesCount++
			p.s.Unlock()
		} else {
			log.Fatalln("invalid nextCall")
		}
		from.s.Lock()
		from.balance -= gasWanted
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
		if mode == 3 {
			// flip-flop calls
			if nextCall == 1 {
				nextCall = 2
			} else {
				nextCall = 1
			}
		} else {
			nextCall = mode
		}
	}
}

func getValuesCount(instance *MiniStore.MiniStore) (uint64, error) {
	res, err := instance.GetArrayDataLength(nil)
	if err != nil {
		return 0, errors.Wrap(err, "fail to GetNumberValue")
	}
	return res.Uint64(), nil
}
