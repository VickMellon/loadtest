package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

var (
	c, n, m, ic uint64
	t, p        time.Duration
	l           bool
	nodes       string
	chainId     string
)

type config struct {
}

func init() {
	flag.Uint64Var(&c, `c`, 1, `Concurrency, number of async threads with requests`)
	flag.Uint64Var(&n, `n`, 0, `Number of transactions to broadcast, 0 - unlimited`)
	flag.Uint64Var(&m, `m`, 0, `Mode: 0 - send Txs, 1 - call SC InsertArray(), 2 - setArrayValues Txs, 3 - setArrayValuesBulk Txs`)
	flag.Uint64Var(&ic, `i`, 1, `Items count: for m=1 - number of values for InsertArray call, for m=2/3 - number of values for setArrayValues(Bulk) Txs`)
	flag.DurationVar(&t, `t`, 0, `Test duration, 0 - unlimited`)
	flag.DurationVar(&p, `p`, 0, `Random delays, up to value`)
	flag.BoolVar(&l, `l`, false, `Logging mode, not interactive`)

	flag.StringVar(&chainId, `chain`, ``, `Chain ID`)
	flag.StringVar(&nodes, `nodes`, ``, `List of REST servers, comma separated (default "http://localhost:8545")`)
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	log.SetOutput(os.Stdout)
}

func main() {
	s := loadState()
	s.checkConfig(chainId, nodes)
	s.requestWorkset(int(c))
	s.updateWallets()
	s.equalizeBalances()
	s.initInstances()
	s.initAuth()
	s.deploySC()
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
	var totalStartBalance uint64
	pr.startBalances = make([]uint64, len(s.workset))
	for i, w := range s.workset {
		pr.startBalances[i] = w.balance
		totalStartBalance += w.balance
		//log.Println("[", w.address, "] ", w.balance)
	}
	var txCost uint64
	if m == 1 {
		txCost = gasWanted * gasPrice
	} else {
		txCost = sendAmount + feesAmount
	}
	log.Println("Initial total balance -", totalStartBalance, ", avg -", int(totalStartBalance)/len(s.workset),
		", estimated Txs -", totalStartBalance/uint64(len(s.workset))/(txCost)*uint64(len(s.workset)))

	// Go!
	for i := 0; i < int(c); i++ {
		time.Sleep(pr.delayUpTo / time.Duration(c)) // initial calls time shift
		if m == 1 {
			go sc_caller(wg, s.workset[i], s.instances[i%len(s.nodes)], pr, ic)
		} else if m >= 2 {
			go rapidIntakeSpender(wg, s.workset[i], s.sc_address, m, ic, s.nodes[i%len(s.nodes)], s.chainId, pr)
		} else {
			go sendTxSpender(wg, s.workset[i], s.workset, sendAmount, s.nodes[i%len(s.nodes)], s.chainId, pr)
		}
		wg.Add(1)
	}
	wg.Wait()
	log.Println("DONE")
}
