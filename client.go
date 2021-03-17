package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	restDefaultBaseUrl = `http://localhost:8545`
	restGetAccount     = `/auth/accounts/`
	restBroadcastTx    = `/txs`
)

const (
	broadcastTxTpl = `{"mode":"%s","tx":%s}`
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
			AccountNumber uint64 `json:"account_number"`
			Sequence      uint64 `json:"sequence"`
		} `json:"value"`
	} `json:"result"`
}

type broadcastResponse struct {
	Txhash string `json:"txhash"`
	Logs   []struct {
		Success bool `json:"success"`
	} `json:"logs"`
}

type account struct {
	address       string
	balance       uint64
	accountNumber uint64
	sequence      uint64
}

var (
	ErrMempoolIsFull    = errors.New("mempool is full")
	ErrTooManyOpenFiles = errors.New("too many open files")
)

func queryAccount(address, nodeUrl string) *account {
	resp, err := http.Get(nodeUrl + restGetAccount + address)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// parse response
	var res accountResponse
	if err = json.Unmarshal(respBody, &res); err != nil {
		log.Println(err)
		return &account{address: address}
	}
	if res.Result.Value.Address != address || len(res.Result.Value.Coins) == 0 {
		return nil
	}
	bal, err := strconv.ParseInt(res.Result.Value.Coins[0].Amount, 10, 64)
	return &account{
		address:       address,
		balance:       uint64(bal),
		accountNumber: res.Result.Value.AccountNumber,
		sequence:      res.Result.Value.Sequence,
	}
}

func broadcastTx(tx string, nodeUrl, mode string) (string, error) {
	btx := fmt.Sprintf(broadcastTxTpl, mode, tx)
	rb := strings.NewReader(btx)
	resp, err := http.Post(nodeUrl+restBroadcastTx, `application/json`, rb)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == http.StatusInternalServerError &&
		strings.Contains(string(respBody), `mempool is full`) {
		return "", ErrMempoolIsFull
	}
	if resp.StatusCode == http.StatusInternalServerError &&
		strings.Contains(string(respBody), `too many open files`) {
		return "", ErrTooManyOpenFiles
	}
	if resp.StatusCode != http.StatusOK ||
		(mode == "sync" && !strings.Contains(string(respBody), `success`)) {
		log.Println("broadcastTx response - ", resp.Status)
		log.Println("broadcastTx response body: ", string(respBody))
		return "", errors.New("broadcastTx error")
	}
	// parse response
	var res broadcastResponse
	if err = json.Unmarshal(respBody, &res); err != nil {
		log.Fatalln(err)
	}
	return res.Txhash, nil
}
