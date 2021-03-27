package main

import (
	"github.com/VickMellon/loadtest/MiniStore"
	"github.com/pkg/errors"
	"log"
	"math/big"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	gasPrice  = uint64(1)
	gasWanted = uint64(50000)
)

var (
	parseSeq = regexp.MustCompile(`[0-9]+`)
)

func sc_caller(wg *sync.WaitGroup, from *wallet, instance *MiniStore.MiniStore, p *process, mode uint64) {
	var err error
	defer wg.Done()
	nextCall := mode
	if mode == 3 {
		nextCall = 1
	}
	retryInt := 10 * time.Millisecond
	for {
		// check source balance
		if from.balance <= gasWanted {
			log.Println("wallet", from.address, " out of tokens:", from.balance)
			return // source wallet out of tokens
		}
		auth := from.auth
		auth.GasLimit = gasWanted
		auth.Nonce = big.NewInt(int64(from.sequence))
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
			if err != nil && parseInstanceError(err) == ErrMempoolIsFull {
				// wait & retry
				log.Println(from.address[:8], "retry after:", retryInt.String())
				time.Sleep(retryInt)
				if retryInt < time.Second {
					retryInt *= 2 // progressive pause, but not longer 1s
				}
				continue
			} else if seq := parseSequenceError(err); seq > 0 {
				// fix failed sequence & retry
				log.Println(from.address[:8], "fix sequence:", from.sequence, " to ", seq)
				from.sequence = seq
				continue
			} else if err != nil {
				log.Println("call SetNumberValue() FAIL, with sequence:", from.sequence, " err:", err)
				return
			}
		} else if nextCall == 2 {
			_, err = instance.AddValue(auth, amount)
			if err != nil && parseInstanceError(err) == ErrMempoolIsFull {
				// wait & retry
				log.Println(from.address[:8], "retry after:", retryInt.String())
				time.Sleep(retryInt)
				if retryInt < time.Second {
					retryInt *= 2 // progressive pause, but not longer 1s
				}
				continue
			} else if seq := parseSequenceError(err); seq > 0 {
				// fix failed sequence & retry
				log.Println(from.address[:8], "fix sequence:", from.sequence, " to ", seq)
				from.sequence = seq
				continue
			} else if err != nil {
				log.Println("call AddValue() FAIL, with sequence:", from.sequence, " err:", err)
				return
			}
			p.s.Lock()
			p.valuesCount++
			p.s.Unlock()
		} else {
			log.Fatalln("invalid nextCall")
		}
		from.balance -= gasWanted
		from.sequence++
		// check process state
		if p.CalcTx() {
			return
		}
		// current delay
		if p.delayUpTo > 0 {
			time.Sleep(p.delayUpTo)
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
		retryInt = 10 * time.Millisecond
	}
}

func getValuesCount(instance *MiniStore.MiniStore) (uint64, error) {
	res, err := instance.GetArrayDataLength(nil)
	if err != nil {
		return 0, errors.Wrap(err, "fail to GetNumberValue")
	}
	return res.Uint64(), nil
}

func parseInstanceError(err error) error {
	if err == nil {
		return nil
	}
	s := err.Error()
	if strings.Contains(s, "-32000") {
		return ErrMempoolIsFull
	}
	return err
}

func parseSequenceError(err error) uint64 {
	//invalid sequence: invalid nonce; got 68, expected 175
	if err == nil {
		return 0
	}
	s := err.Error()
	if strings.Contains(s, "invalid sequence") {
		res := parseSeq.FindAllString(s, -1)
		if len(res) == 2 {
			seq, err := strconv.ParseInt(res[1], 10, 64)
			if err == nil {
				return uint64(seq)
			}
		}
	}
	return 0
}
