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
		requestAt := time.Now()
		if nextCall == 1 {
			_, err = instance.SetNumberValue(auth, amount)
			if err = parseInstanceError(err); err == ErrMempoolIsFull || err == ErrTooManyOpenFiles {
				// wait & retry
				//log.Println(from.address[:8], "retry after:", retryInt.String())
				time.Sleep(retryInt)
				if retryInt < 100*time.Millisecond {
					retryInt *= 2 // progressive pause, but not longer 1s
				}
				continue
			} else if seq := parseSequenceError(err); seq > 0 {
				// fix failed sequence & retry
				//if from.sequence > seq {
				//	time.Sleep(retryInt)
				//	if retryInt < 100*time.Millisecond {
				//		retryInt *= 2 // progressive pause, but not longer 1s
				//	}
				//	continue
				//}
				log.Println(from.address[:8], "fix sequence:", from.sequence, " to ", seq)
				from.sequence = seq
				continue
			} else if err != nil {
				log.Println("call SetNumberValue() FAIL, with sequence:", from.sequence, " err:", err)
				return
			}
		} else if nextCall == 2 {
			idx := big.NewInt(rand.Int63n(32768)) // total array length up to 32K
			_, err = instance.SetArrayValue(auth, idx, amount)
			if err = parseInstanceError(err); err == ErrMempoolIsFull || err == ErrTooManyOpenFiles {
				// wait & retry
				//log.Println(from.address[:8], "retry after:", retryInt.String())
				time.Sleep(retryInt)
				if retryInt < 100*time.Millisecond {
					retryInt *= 2 // progressive pause, but not longer 100ms
				}
				continue
			} else if seq := parseSequenceError(err); seq > 0 {
				// fix failed sequence & retry
				//if from.sequence > seq {
				//	time.Sleep(retryInt)
				//	if retryInt < 100*time.Millisecond {
				//		retryInt *= 2 // progressive pause, but not longer 1s
				//	}
				//	continue
				//}
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
		// delay consider previous request duration
		if p.delayUpTo > 0 {
			timePassed := time.Now().Sub(requestAt)
			if timePassed < p.delayUpTo {
				time.Sleep(p.delayUpTo - timePassed)
			}
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
	//log.Println(s)
	if strings.Contains(s, "-32000") {
		return ErrMempoolIsFull
	}
	if strings.Contains(s, "too many open files") {
		return ErrTooManyOpenFiles
	}
	if strings.Contains(s, "EOF") {
		// assume request timeout but Tx was passed..
		return nil
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
