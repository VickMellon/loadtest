package main

import (
	"encoding/base64"
	"encoding/json"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"io/ioutil"
	"log"
	"os"
)

const (
	stateFileName = `loadtest_state.json`
)

type state struct {
	wallets []*wallet
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

type wallet struct {
	privKey       secp256k1.PrivKeySecp256k1
	pubKey        secp256k1.PubKeySecp256k1
	name          string // contract: also used as password
	address       string // bech32
	balance       uint64 // uatolo
	accountNumber uint64
	sequence      uint64
}

type storageState struct {
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

func saveState(s *state) {
	var ss storageState
	ss.Wallets = make([]storageWallet, 0, len(s.wallets))
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
