package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os/exec"
	"strconv"
	"time"
)

const (
	nodeFile = `nodef`
	cliFile  = `clif`

	restBaseUrl    = `http://localhost:1317`
	restGetAccount = restBaseUrl + `/auth/accounts/`
)

type accountResponse struct {
	Result struct {
		Type  string `json:"type"`
		Value struct {
			Address string `json:"address"`
			Coins   []struct {
				Denom  string `json:"denom"`
				Amount string `json:"amount"`
			} `json:"coins"`
			AccountNumber string `json:"account_number"`
			Sequence      string `json:"sequence"`
		} `json:"value"`
	} `json:"result"`
}

type account struct {
	Address       string
	Balance       uint64
	AccountNumber uint64
	Sequence      uint64
}

func queryAccount(address string) *account {
	t1 := time.Now()
	code, respBody, err := fasthttp.Get(nil, restGetAccount+address)
	if err != nil {
		log.Fatal(err)
	}
	if code != fasthttp.StatusOK {
		log.Fatal("Response code not OK", code)
	}
	// parse response
	var res accountResponse
	if err = json.Unmarshal(respBody, &res); err != nil {
		log.Fatal(err)
	}
	if res.Result.Value.Address != address {
		return nil
	}
	bal, err := strconv.ParseInt(res.Result.Value.Coins[0].Amount, 10, 64)
	anum, err := strconv.ParseInt(res.Result.Value.AccountNumber, 10, 64)
	seq, err := strconv.ParseInt(res.Result.Value.Sequence, 10, 64)
	fmt.Println(time.Now().Sub(t1))
	return &account{
		Address:       address,
		Balance:       uint64(bal),
		AccountNumber: uint64(anum),
		Sequence:      uint64(seq),
	}
}

func getKeysList() map[string]localKey {
	c := exec.Command(cliFile, `keys`, `list`)
	out, err := c.Output()
	if err != nil {
		log.Fatal(err)
	}
	var keys []localKey
	if err = json.Unmarshal(out, &keys); err != nil {
		log.Fatal(err)
	}
	res := make(map[string]localKey, len(keys))
	for _, k := range keys {
		res[k.Name] = k
	}
	return res
}

func addNewKey(name string) localKey {
	// check for existed
	if k, ok := keys[name]; ok {
		return k
	}
	c := exec.Command(cliFile, `keys`, `add`, name)
	// prepare input
	var inb bytes.Buffer
	inb.WriteString(name + "\n") // input password (let it be the same as name)
	inb.WriteString(name + "\n") // repeat password
	c.Stdin = &inb
	// there is bug - clif output key data to StdErr instead of StdOut (already fixed in latest CosmosSDK)
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	var key localKey
	if err = json.Unmarshal(out, &key); err != nil {
		log.Fatal(err)
	}
	return key
}
