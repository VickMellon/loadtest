package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	stateFileName = `loadtest_state.json`
)

type state struct {
	chainId string
	nodes   []string
	wallets []*wallet
	workset []*wallet // specified part of all wallets for current run
}

type wallet struct {
	privKey       secp256k1.PrivKeySecp256k1
	pubKey        secp256k1.PubKeySecp256k1
	name          string // contract: also used as password
	address       string // bech32
	balance       uint64 // uatolo
	accountNumber uint64
	sequence      uint64
	s             sync.Mutex // for async access
}

type storageState struct {
	ChainId string          `json:"chain_id"`
	Nodes   []string        `json:"nodes"`
	Wallets []storageWallet `json:"wallets"`
}

type storageWallet struct {
	Name string `json:"name"` // contract: also used as password
	Pk   string `json:"pk"`   // base64 encoded
}

func loadState() (s *state) {
	var ok bool
	s = &state{}
	d, err := ioutil.ReadFile(stateFileName)
	if err != nil {
		log.Println("Can't read state file", err)
		return
	}
	var ss storageState
	if err = json.Unmarshal(d, &ss); err != nil {
		log.Println("Can't parse state file", err)
		return
	}
	s.chainId = ss.ChainId
	s.nodes = ss.Nodes
	s.wallets = make([]*wallet, 0, len(ss.Wallets))
	for _, sw := range ss.Wallets {
		w := &wallet{
			name: sw.Name,
		}
		pk, err := base64.StdEncoding.DecodeString(sw.Pk)
		if err != nil {
			log.Println("Can't decode private key", err)
			continue
		}
		if len(pk) != 32 {
			log.Println("Invalid private key len", err)
			continue
		}
		w.privKey = secp256k1.PrivKeySecp256k1{}
		copy(w.privKey[:], pk)

		if w.pubKey, ok = w.privKey.PubKey().(secp256k1.PubKeySecp256k1); !ok {
			log.Println("pubKey is not secp256k1.PubKeySecp256k1")
			continue
		}
		w.address, err = addressFromPubKey(w.pubKey)
		if err != nil {
			log.Println("Can't ger Friday address from public key", err)
			continue
		}
		s.wallets = append(s.wallets, w)
	}
	return
}

func (s *state) checkConfig(chainId string, nodes string) {
	var isUpdated bool
	if chainId != "" {
		s.chainId = chainId
		isUpdated = true
	}
	if s.chainId == "" {
		log.Fatalln("chainId is empty, please, specify --chain parameter")
	}
	if nodes != "" {
		s.nodes = strings.Split(nodes, `,`)
		isUpdated = true
	}
	if len(s.nodes) == 0 {
		s.nodes = []string{restDefaultBaseUrl}
		log.Println("nodes URL list is empty, default address will be used", restDefaultBaseUrl)
	}
	for _, r := range s.nodes {
		if u, err := url.Parse(r); err != nil || u == nil {
			log.Fatalln("node URL is invalid")
		} else if u.Scheme == "" {
			log.Fatalln("node URL is invalid, scheme required")
		} else if u.Port() == "" {
			log.Fatalln("node URL is invalid, port required")
		}
	}
	if isUpdated {
		s.saveState()
	}
}

func (s *state) newWallet(name string) {
	pk := newPrivKey()
	pb, _ := pk.PubKey().(secp256k1.PubKeySecp256k1)
	addr, _ := addressFromPubKey(pb)
	s.wallets = append(s.wallets, &wallet{
		privKey:       pk,
		pubKey:        pb,
		name:          name,
		address:       addr,
		balance:       0,
		accountNumber: 0,
		sequence:      0,
	})
}

func (s *state) saveState() {
	ss := storageState{
		ChainId: s.chainId,
		Nodes:   s.nodes,
		Wallets: make([]storageWallet, 0, len(s.wallets)),
	}
	for _, w := range s.wallets {
		ss.Wallets = append(ss.Wallets, storageWallet{
			Name: w.name,
			Pk:   base64.StdEncoding.EncodeToString(w.privKey[:]),
		})
	}
	stateFile, err := os.Create(stateFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer stateFile.Close()
	encoder := json.NewEncoder(stateFile)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(ss)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *state) requestWorkset(n int) {
	if n == 1 {
		n = 2 // minimal workset is 2 wallets
	}
	if n > len(s.wallets) {
		// generate more wallets
		for i := len(s.wallets) + 1; i <= n; i++ {
			s.newWallet(fmt.Sprintf("loadtest%d", i))
		}
		s.saveState()
	}
	// limit workset with only requested count of wallets
	s.workset = s.wallets[:n]
}

func (s *state) equalizeBalances() {
	n := len(s.workset)
	m := len(s.wallets)
	if n == 0 {
		log.Fatalln("empty workset, requestWorkset first")
	} else if n == 1 {
		// nothing to do for now
		return
	}
	bank := s.workset[0]
	var total uint64
	for _, w := range s.wallets {
		total += w.balance
	}
	if total <= feesAmount*uint64(m+n-2) {
		log.Fatalln("total balance is too low, charge more tokens to", bank.address)
	}
	total -= feesAmount * uint64(m+n-2) // exclude balancing fees
	each := total / uint64(n)           // target balance of each wallet in workset
	if each <= 2*feesAmount {
		log.Fatalln("total balance too low, charge more tokens to", bank.address)
	}
	var txs uint64
	// stage 1 - cut excess balances to bank from all wallets
	for i := 1; i < m; i++ {
		amount := uint64(0)
		if i < n && s.wallets[i].balance > each+feesAmount {
			// this balance too high - transfer difference to bank
			amount = s.wallets[i].balance - (each + feesAmount)
		} else if i >= n && s.wallets[i].balance > feesAmount {
			// some balance on wallet out of workset - move entire sum to bank
			amount = s.wallets[i].balance - feesAmount
		}
		if amount > 0 {
			tx := getSignedSendTx(s.wallets[i].address, bank.address, amount,
				"equalizeBalances", s.wallets[i].privKey, s.chainId, s.wallets[i].accountNumber, s.wallets[i].sequence)
			_, err := broadcastTx(tx, s.nodes[0], "sync")
			for err == ErrMempoolIsFull {
				// wait & retry
				time.Sleep(100 * time.Millisecond)
				_, err = broadcastTx(tx, s.nodes[0], "sync")
			}
			if err != nil {
				log.Fatalln("equalizeBalances failed,", err)
			}
			s.wallets[i].balance -= amount + feesAmount
			s.wallets[i].sequence++
			bank.balance += amount
			txs++
			time.Sleep(2 * time.Millisecond)
		}
	}
	// stage 2 - deposit low balances from bank for workset wallets only
	for i := 1; i < n; i++ {
		if s.workset[i].balance < each-feesAmount {
			// this balance too low - get deposit from bank
			dif := each - s.workset[i].balance
			tx := getSignedSendTx(bank.address, s.workset[i].address, dif,
				"equalizeBalances", bank.privKey, s.chainId, bank.accountNumber, bank.sequence)
			_, err := broadcastTx(tx, s.nodes[0], "sync")
			for err == ErrMempoolIsFull {
				// wait & retry
				time.Sleep(100 * time.Millisecond)
				_, err = broadcastTx(tx, s.nodes[0], "sync")
			}
			if err != nil {
				log.Fatalln("equalizeBalances failed,", err)
			}
			bank.balance -= dif + feesAmount
			bank.sequence++
			s.workset[i].balance += dif
			txs++
			time.Sleep(2 * time.Millisecond)
		}
	}
	if txs > 0 {
		log.Println("waiting 10s for sure commits after equalize balances")
		time.Sleep(time.Second * 10)
	}
}

func (s *state) updateWorkset() {
	if len(s.workset) == 0 {
		log.Fatalln("updateWorkset failed - workset is empty")
	}
	updateW(s.workset, s.nodes[0])
}
func (s *state) updateWallets() {
	if len(s.wallets) == 0 {
		log.Fatalln("updateWallets failed - wallets is empty")
	}
	updateW(s.wallets, s.nodes[0])
}
func updateW(wallets []*wallet, baseUrl string) {
	wg := &sync.WaitGroup{}
	for i := range wallets {
		wg.Add(1)
		go func(k int, sw []*wallet) {
			a := queryAccount(sw[k].address, baseUrl)
			if a != nil {
				sw[k].balance = a.balance
				sw[k].accountNumber = a.accountNumber
				sw[k].sequence = a.sequence
			} else {
				sw[k].balance = 0
				sw[k].accountNumber = 0
				sw[k].sequence = 0
			}
			wg.Done()
		}(i, wallets)
		time.Sleep(2 * time.Millisecond) // to keep friendly rps rate
	}
	wg.Wait()
}
