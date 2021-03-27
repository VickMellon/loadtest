package main

import (
	"flag"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	c, n, m uint64
	t, p    time.Duration
	l       bool
	nodes   string
	chainId string
)

type config struct {
}

func init() {
	flag.Uint64Var(&c, `c`, 1, `Concurrency, number of async threads with requests`)
	flag.Uint64Var(&n, `n`, 0, `Number of transactions to broadcast, 0 - unlimited`)
	flag.Uint64Var(&m, `m`, 0, `Mode: 0 - send Txs, 1 - call SetNumberValue, 2 - call AddValue, 3 - call both SetNumberValue and AddValue`)
	flag.DurationVar(&t, `t`, 0, `Test duration, 0 - unlimited`)
	flag.DurationVar(&p, `p`, 0, `Random delays, up to value`)
	flag.BoolVar(&l, `l`, false, `Logging mode, not interactive`)

	flag.StringVar(&chainId, `chain`, ``, `Chain ID`)
	flag.StringVar(&nodes, `nodes`, ``, `List of REST servers, comma separated (default "http://localhost:8545")`)
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
}

func main() {
	s := loadState()
	s.checkConfig(chainId, nodes)
	s.requestWorkset(int(c))
	s.updateWallets()
	s.initInstances()
	s.initAuth()
	s.deploySC()
	s.equalizeBalances()
	s.updateWorkset()

	wg := &sync.WaitGroup{}
	pr := &process{
		startedAt:   time.Now(),
		txLimit:     n,
		delayUpTo:   p,
		delayNow:    p / 2,
		interactive: !l,
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
	var txCost uint64
	if m > 0 {
		txCost = gasWanted * gasPrice
	} else {
		txCost = sendAmount + feesAmount
	}
	log.Println("Initial total balance -", totalStartBalance, ", avg -", int(totalStartBalance)/len(s.workset),
		", estimated Txs -", totalStartBalance/uint64(len(s.workset))/(txCost)*uint64(len(s.workset)))
	if m > 1 {
		var err error
		if pr.valuesCount, err = getValuesCount(s.instances[0]); err != nil {
			log.Fatalln("Can't get initial SC array values count:", err)
		}
		log.Println("Initial SC array values count:", pr.valuesCount)
	}
	// Go!
	for i := 0; i < int(c); i++ {
		if m > 0 {
			log.Println("All SC calls from wallet", s.workset[i].address, " will be sent to node:", s.nodes[i%len(s.nodes)])
			go sc_caller(wg, s.workset[i], s.instances[i%len(s.nodes)], pr, m)
		} else {
			go spender(wg, s.workset[i], s.workset, sendAmount, s.nodes[i%len(s.nodes)], s.chainId, pr)
		}
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
		", estimated Txs -", totalFinishBalance/uint64(len(s.workset))/(txCost)*uint64(len(s.workset)))
	{
		log.Println("Checking final SC array values count...")
		actual, err := getValuesCount(s.instances[0])
		if err != nil {
			log.Fatalln("Can't get final SC array values count:", err)
		}
		if actual == pr.valuesCount {
			log.Println("Final SC array values count:", pr.valuesCount, " - MATCHED")
		} else {
			log.Println("FAIL! expected final SC array values count not equal to actual: ", pr.valuesCount, "!=", actual)
		}
	}
}
