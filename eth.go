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
	gasPrice              = uint64(1)
	gasWanted             = uint64(50000)
	gasWantedPerArrayItem = uint64(7000)
)

var (
	parseSeq = regexp.MustCompile(`[0-9]+`)
)

func sc_caller(wg *sync.WaitGroup, from *wallet, instance *MiniStore.MiniStore, p *process, ic uint64) {
	var err error
	defer wg.Done()
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
		requestAt := time.Now()
		auth.GasLimit = gasWanted + gasWantedPerArrayItem*ic
		idx := big.NewInt(rand.Int63n(32768)) // total array length up to 32K
		amounts := make([]*big.Int, ic)
		for i := range amounts {
			amounts[i] = big.NewInt(rand.Int63())
		}
		_, err = instance.InsertArray(auth, idx, amounts)
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
			log.Println(from.address[:8], "fix sequence:", from.sequence, " to ", seq)
			from.sequence = seq
			continue
		} else if err != nil && strings.Contains(err.Error(), "internal") {
			log.Println("call InsertArray() FAIL, err:", err, ", retrying within", retryInt)
			time.Sleep(retryInt)
			if retryInt < 30*time.Second {
				retryInt *= 2 // progressive pause, but not longer 30s
			}
			continue
		} else if err != nil {
			log.Println("call InsertArray() FAIL, with sequence:", from.sequence, " err:", err)
			return
		}
		p.s.Lock()
		p.valuesCount++
		p.s.Unlock()

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
