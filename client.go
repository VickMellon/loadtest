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
	restPort           = `:1317`
	ethPort            = `:8545`
	restDefaultBaseUrl = `http://localhost`
	restGetAccount     = `/cosmos/auth/v1beta1/accounts/`
	restGetBalance     = `/cosmos/bank/v1beta1/balances/`
	restBroadcastTx    = `/txs`
)

const (
	broadcastTxTpl = `{"mode":"%s","tx":%s}`
)

type accountResponse struct {
	Account struct {
		Type        string `json:"type"`
		BaseAccount struct {
			Address       string      `json:"address"`
			AccountNumber interface{} `json:"account_number"`
			Sequence      interface{} `json:"sequence"`
		} `json:"base_account"`
	} `json:"account"`
}

type balanceResponse struct {
	Balances []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balances"`
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
	resp, err := http.Get(nodeUrl + restPort + restGetAccount + address)
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
		log.Println("failed to parse accountResponse:", err)
		return &account{address: address}
	}
	if res.Account.BaseAccount.Address != address {
		return &account{address: address}
	}
	var accountNumber uint64
	if s, ok := res.Account.BaseAccount.AccountNumber.(string); ok {
		v, _ := strconv.ParseInt(s, 10, 64)
		accountNumber = uint64(v)
	} else if v, ok := res.Account.BaseAccount.AccountNumber.(float64); ok {
		accountNumber = uint64(v)
	}
	var sequence uint64
	if s, ok := res.Account.BaseAccount.Sequence.(string); ok {
		v, _ := strconv.ParseInt(s, 10, 64)
		sequence = uint64(v)
	} else if v, ok := res.Account.BaseAccount.Sequence.(float64); ok {
		sequence = uint64(v)
	}
	// get balance
	resp, err = http.Get(nodeUrl + restPort + restGetBalance + address)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	respBody, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	// parse response
	var res2 balanceResponse
	if err = json.Unmarshal(respBody, &res2); err != nil {
		log.Println("failed to parse balanceResponse:", err)
		return &account{address: address}
	}
	var bal int64
	for _, c := range res2.Balances {
		if c.Denom == "aphoton" {
			bal, _ = strconv.ParseInt(res2.Balances[0].Amount, 10, 64)
		}
	}
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
	resp, err := http.Post(nodeUrl+restPort+restBroadcastTx, `application/json`, rb)
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
