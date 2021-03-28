package main

import (
	"encoding/base64"
	"fmt"
	"github.com/cosmos/ethermint/crypto/ethsecp256k1"
	"github.com/tendermint/tendermint/crypto"
	"log"
)

const (
	sendAmount = uint64(1)
	feesAmount = uint64(0)
	toSigTpl   = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"%d","denom":"aphoton"}],"gas":"200000"},"memo":"%s","msgs":[{"type":"cosmos-sdk/MsgSend","value":{"amount":[{"amount":"%d","denom":"aphoton"}],"from_address":"%s","to_address":"%s"}}],"sequence":"%d"}`
	sigTpl     = `[{"pub_key":{"type":"ethermint/PubKeyEthSecp256k1","value":"%s"},"signature":"%s"}]`
	sendTxTpl  = `{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"%s","to_address":"%s","amount":[{"denom":"aphoton","amount":"%d"}]}}],"fee":{"amount":[{"denom":"aphoton","amount":"%d"}],"gas":"200000"},"signatures":%s,"memo":"%s"}`
)

func getSignedSendTx(from, to string, amount uint64, memo string,
	pk crypto.PrivKey, chainId string, accountNumber, sequence uint64) string {

	sigDoc := fmt.Sprintf(toSigTpl, accountNumber, chainId, feesAmount, memo, amount, from, to, sequence)
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
