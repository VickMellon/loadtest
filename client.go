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
			AccountNumber interface{} `json:"account_number"`
			Sequence      interface{} `json:"sequence"`
		} `json:"value"`
	} `json:"result"`
}

type broadcastResponse struct {
	Txhash string `json:"txhash"`
	Logs   []struct {
		Success bool `json:"success"`
	} `json:"logs"`
	Code uint64 `json:"code"`
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
	ErrSequenceWrong    = errors.New("wrong sequence")
	ErrEOF              = errors.New("EOF")
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
	var accountNumber uint64
	if s, ok := res.Result.Value.AccountNumber.(string); ok {
		v, _ := strconv.ParseInt(s, 10, 64)
		accountNumber = uint64(v)
	} else if v, ok := res.Result.Value.AccountNumber.(float64); ok {
		accountNumber = uint64(v)
	}
	var sequence uint64
	if s, ok := res.Result.Value.Sequence.(string); ok {
		v, _ := strconv.ParseInt(s, 10, 64)
		sequence = uint64(v)
	} else if v, ok := res.Result.Value.Sequence.(float64); ok {
		sequence = uint64(v)
	}
	bal, err := strconv.ParseInt(res.Result.Value.Coins[0].Amount, 10, 64)
	return &account{
		address:       address,
		balance:       uint64(bal),
		accountNumber: accountNumber,
		sequence:      sequence,
	}
}

func broadcastTx(tx string, nodeUrl, mode string) (string, error) {
	btx := fmt.Sprintf(broadcastTxTpl, mode, tx)
	rb := strings.NewReader(btx)
	resp, err := http.Post(nodeUrl+restBroadcastTx, `application/json`, rb)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	//log.Println(resp.StatusCode, string(respBody))
	if resp.StatusCode == http.StatusInternalServerError &&
		strings.Contains(string(respBody), `mempool is full`) {
		return "", ErrMempoolIsFull
	}
	if resp.StatusCode == http.StatusInternalServerError &&
		strings.Contains(string(respBody), `too many open files`) {
		return "", ErrTooManyOpenFiles
	}
	if resp.StatusCode != http.StatusOK {
		log.Println("broadcastTx response - ", resp.Status)
		log.Println("broadcastTx response body: ", string(respBody))
		return "", errors.New("broadcastTx error")
	}
	// parse response
	var res broadcastResponse
	if err = json.Unmarshal(respBody, &res); err != nil {
		log.Fatalln(err)
	}
	if res.Code > 0 {
		if res.Code == 20 {
			return "", ErrMempoolIsFull
		}
		if res.Code == 4 {
			return "", ErrSequenceWrong
		}
		log.Println(resp.StatusCode, string(respBody))
		return "", errors.New("broadcastTx error")
	}
	return res.Txhash, nil
}
