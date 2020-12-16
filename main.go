package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

var (
	c, n    uint64
	t       time.Duration
	nodes   string
	chainId string
)

type config struct {
}

func init() {
	flag.Uint64Var(&c, `c`, 1, `Concurrency, number of async threads with requests`)
	flag.Uint64Var(&n, `n`, 0, `Number of transactions to broadcast, 0 - unlimited`)
	flag.DurationVar(&t, `t`, 0, `Test duration, 0 - unlimited`)

	flag.StringVar(&chainId, `chain`, ``, `Chain ID`)
	flag.StringVar(&nodes, `nodes`, ``, `List of REST servers, comma separated (default "http://localhost:1317")`)
	flag.Parse()
}

func main() {
	s := loadState()
	s.checkConfig(chainId, nodes)
	s.requestWorkset(int(c))
	s.updateWallets()
	s.equalizeBalances()
	s.updateWorkset()

	wg := &sync.WaitGroup{}
	pr := &process{
		startedAt: time.Now(),
		txLimit:   n,
	}
	if t > 0 {
		pr.mustStopAfter = pr.startedAt.Add(t)
	}
	var totalStartBalance, totalFinishBalance uint64
	pr.startBalances = make([]uint64, len(s.workset))
	for i, w := range s.workset {
		pr.startBalances[i] = w.balance
		totalStartBalance += w.balance
		//log.Println("[", w.address, "] ", w.balance)
	}
	log.Println("Initial total balance -", totalStartBalance, ", avg -", int(totalStartBalance)/len(s.workset),
		", estimated Txs -", totalStartBalance/uint64(len(s.workset))/(sendAmount+feesAmount)*uint64(len(s.workset)))
	// Go!
	for i := 0; i < int(c); i++ {
		go spender(wg, s.workset[i], s.workset, sendAmount, s.nodes[i%len(s.nodes)], s.chainId, pr)
		wg.Add(1)
	}
	wg.Wait()
	log.Println("waiting 30s for sure commits")
	time.Sleep(time.Second * 30)
	pr.finishBalances = make([]uint64, len(s.workset))
	for i, w := range s.workset {
		pr.finishBalances[i] = w.balance
		totalFinishBalance += w.balance
	}
	s.updateWorkset()
	log.Println("Checking final balances...")
	ok := true
	for i, w := range s.workset {
		if pr.finishBalances[i] != w.balance {
			ok = false
			log.Println("FAIL! expected final balance not equal to actual: [", w.address, "] ", pr.finishBalances[i], "!=", w.balance)
		}
		//log.Println("[", w.address, "] ", w.balance)
	}
	if ok {
		log.Println("SUCCESS")
	}
	log.Println("Final total balance -", totalFinishBalance, ", avg -", int(totalFinishBalance)/len(s.workset),
		", estimated Txs -", totalFinishBalance/uint64(len(s.workset))/(sendAmount+feesAmount)*uint64(len(s.workset)))
}
