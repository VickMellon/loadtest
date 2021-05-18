package main

import (
	"encoding/base64"
	"fmt"
	"github.com/cosmos/ethermint/crypto/ethsecp256k1"
	"github.com/ethereum/go-ethereum/common"
	"github.com/tendermint/tendermint/crypto"
	"log"
	"strings"
)

const (
	sendAmount              = uint64(1)
	feesAmount              = uint64(0)
	gasWantedPerArrayItemRI = 5000

	sigTpl = `[{"pub_key":{"type":"ethermint/PubKeyEthSecp256k1","value":"%s"},"signature":"%s"}]`

	toSigSendTxTpl = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"%d","denom":"aphoton"}],"gas":"200000"},"memo":"%s","msgs":[{"type":"cosmos-sdk/MsgSend","value":{"amount":[{"amount":"%d","denom":"aphoton"}],"from_address":"%s","to_address":"%s"}}],"sequence":"%d"}`
	sendTxTpl      = `{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"%s","to_address":"%s","amount":[{"denom":"aphoton","amount":"%d"}]}}],"fee":{"amount":[{"denom":"aphoton","amount":"%d"}],"gas":"200000"},"signatures":%s,"memo":"%s"}`

	toSigSetValueTxTpl = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"%d","denom":"aphoton"}],"gas":"200000"},"memo":"%s","msgs":[{"type":"rapidintake/MsgSetValue","value":{"from_address":"%s","sc_address":"%s","slot_num":"%s","value":"%s"}}],"sequence":"%d"}`
	setValueTxTpl      = `{"msg":[{"type":"rapidintake/MsgSetValue","value":{"from_address":"%s","sc_address":"%s","slot_num":"%s","value":"%s"}}],"fee":{"amount":[{"denom":"aphoton","amount":"%d"}],"gas":"200000"},"signatures":%s,"memo":"%s"}`

	toSigAddArrayValueTxTpl = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"%d","denom":"aphoton"}],"gas":"200000"},"memo":"%s","msgs":[{"type":"rapidintake/MsgAddArrayValue","value":{"from_address":"%s","sc_address":"%s","slot_num":"%s","value":"%s"}}],"sequence":"%d"}`
	addArrayValueTxTpl      = `{"msg":[{"type":"rapidintake/MsgAddArrayValue","value":{"from_address":"%s","sc_address":"%s","slot_num":"%s","value":"%s"}}],"fee":{"amount":[{"denom":"aphoton","amount":"%d"}],"gas":"200000"},"signatures":%s,"memo":"%s"}`

	toSigSetArrayValuesTxTpl = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"%d","denom":"aphoton"}],"gas":"%d"},"memo":"%s","msgs":[{"type":"rapidintake/MsgSetArrayValues","value":{"from_address":"%s","index":"%s","sc_address":"%s","slot_num":"%s","values":[%s]}}],"sequence":"%d"}`
	setArrayValueTxTpl       = `{"msg":[{"type":"rapidintake/MsgSetArrayValues","value":{"from_address":"%s","index":"%s","sc_address":"%s","slot_num":"%s","values":[%s]}}],"fee":{"amount":[{"denom":"aphoton","amount":"%d"}],"gas":"%d"},"signatures":%s,"memo":"%s"}`
)

func getSignedSendTx(from, to string, amount uint64, memo string,
	pk crypto.PrivKey, chainId string, accountNumber, sequence uint64) string {

	sigDoc := fmt.Sprintf(toSigSendTxTpl, accountNumber, chainId, feesAmount, memo, amount, from, to, sequence)
	sig, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	pb, ok := pk.PubKey().(ethsecp256k1.PubKey)
	if !ok {
		log.Fatal("not ethsecp256k1.PubKey")
	}
	pubStr := base64.StdEncoding.EncodeToString(pb[:])
	sigBody := fmt.Sprintf(sigTpl, pubStr, sigStr)

	return fmt.Sprintf(sendTxTpl, from, to, amount, feesAmount, sigBody, memo)
}

func getSignedSetValueTx(from string, to common.Address, slot, value common.Hash, memo string,
	pk crypto.PrivKey, chainId string, accountNumber, sequence uint64) string {

	toEnc := base64.StdEncoding.EncodeToString(to.Bytes())
	slotEnc := base64.StdEncoding.EncodeToString(slot.Bytes())
	valEnc := base64.StdEncoding.EncodeToString(value.Bytes())
	sigDoc := fmt.Sprintf(toSigSetValueTxTpl, accountNumber, chainId, feesAmount, memo, from, toEnc, slotEnc, valEnc, sequence)
	sig, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	pb, ok := pk.PubKey().(ethsecp256k1.PubKey)
	if !ok {
		log.Fatal("not ethsecp256k1.PubKey")
	}
	pubStr := base64.StdEncoding.EncodeToString(pb[:])
	sigBody := fmt.Sprintf(sigTpl, pubStr, sigStr)

	return fmt.Sprintf(setValueTxTpl, from, to.String(), slot.String(), value.String(), feesAmount, sigBody, memo)
}

func getSignedAddArrayValueTx(from string, to common.Address, slot, value common.Hash, memo string,
	pk crypto.PrivKey, chainId string, accountNumber, sequence uint64) string {

	toEnc := base64.StdEncoding.EncodeToString(to.Bytes())
	slotEnc := base64.StdEncoding.EncodeToString(slot.Bytes())
	valEnc := base64.StdEncoding.EncodeToString(value.Bytes())
	sigDoc := fmt.Sprintf(toSigAddArrayValueTxTpl, accountNumber, chainId, feesAmount, memo, from, toEnc, slotEnc, valEnc, sequence)
	sig, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	pb, ok := pk.PubKey().(ethsecp256k1.PubKey)
	if !ok {
		log.Fatal("not ethsecp256k1.PubKey")
	}
	pubStr := base64.StdEncoding.EncodeToString(pb[:])
	sigBody := fmt.Sprintf(sigTpl, pubStr, sigStr)

	return fmt.Sprintf(addArrayValueTxTpl, from, to.String(), slot.String(), value.String(), feesAmount, sigBody, memo)
}

func getSignedSetArrayValuesTx(from string, to common.Address, slot, index common.Hash, values []common.Hash, memo string,
	pk crypto.PrivKey, chainId string, accountNumber, sequence uint64) string {

	toEnc := base64.StdEncoding.EncodeToString(to.Bytes())
	slotEnc := base64.StdEncoding.EncodeToString(slot.Bytes())
	indexEnc := base64.StdEncoding.EncodeToString(index.Bytes())
	valsEnc := make([]string, len(values))
	for i, v := range values {
		valsEnc[i] = `"` + base64.StdEncoding.EncodeToString(v.Bytes()) + `"`
	}
	vals := strings.Join(valsEnc, `,`)
	gasWanted := gasWantedPerArrayItemRI*len(values) + 200000
	sigDoc := fmt.Sprintf(toSigSetArrayValuesTxTpl,
		accountNumber, chainId, feesAmount, gasWanted, memo,
		from, indexEnc, toEnc, slotEnc, vals, sequence)
	sig, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	pb, ok := pk.PubKey().(ethsecp256k1.PubKey)
	if !ok {
		log.Fatal("not ethsecp256k1.PubKey")
	}
	pubStr := base64.StdEncoding.EncodeToString(pb[:])
	sigBody := fmt.Sprintf(sigTpl, pubStr, sigStr)
	for i, v := range values {
		valsEnc[i] = `"` + v.String() + `"`
	}
	vals = strings.Join(valsEnc, `,`)
	return fmt.Sprintf(setArrayValueTxTpl, from, index.String(), to.String(), slot.String(), vals, feesAmount, gasWanted, sigBody, memo)
}
