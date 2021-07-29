package main

import (
	"encoding/base64"
	"fmt"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/ethereum/go-ethereum/common"
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
	//sendTxTpl      = `{"body":{"messages":[{"@type":"/cosmos.bank.v1beta1.MsgSend","from_address":"eth1hdstgqw0dan6yqg3swjwmcnhqv680cyfuknlcf","to_address":"eth1887l3eh0znazfaavm68n980urh3efytr0zdc7x","amount":[{"denom":"aphoton","amount":"10000000000000"}]}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}`

	toSigSetArrayValuesTxTpl = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"%d","denom":"aphoton"}],"gas":"%d"},"memo":"%s","msgs":[{"type":"rapidintake/%s","value":{"from_address":"%s","index":"%s","sc_address":"%s","slot_num":"%s","values":[%s]}}],"sequence":"%d"}`
	setArrayValuesTxTpl      = `{"msg":[{"type":"rapidintake/%s","value":{"from_address":"%s","index":"%s","sc_address":"%s","slot_num":"%s","values":[%s]}}],"fee":{"amount":[{"denom":"aphoton","amount":"%d"}],"gas":"%d"},"signatures":%s,"memo":"%s"}`
)

func getSignedSendTx(from, to string, amount uint64, memo string,
	pk cryptotypes.PrivKey, chainId string, accountNumber, sequence uint64) string {

	sigDoc := fmt.Sprintf(toSigSendTxTpl, accountNumber, chainId, feesAmount, memo, amount, from, to, sequence)
	sig, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	pb := pk.PubKey()
	pubStr := base64.StdEncoding.EncodeToString(pb.Bytes()[:])
	sigBody := fmt.Sprintf(sigTpl, pubStr, sigStr)

	return fmt.Sprintf(sendTxTpl, from, to, amount, feesAmount, sigBody, memo)
}

func getSignedSetArrayValuesTx(from string, to common.Address, slot, index common.Hash, values []common.Hash, bulk bool,
	pk cryptotypes.PrivKey, chainId string, accountNumber, sequence uint64) string {
	msgName := "MsgSetArrayValues"
	if bulk {
		msgName = "MsgSetArrayValuesBulk"
	}
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
		accountNumber, chainId, feesAmount, gasWanted, "", msgName,
		from, indexEnc, toEnc, slotEnc, vals, sequence)
	sig, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	sigStr := base64.StdEncoding.EncodeToString(sig)
	pb := pk.PubKey()
	pubStr := base64.StdEncoding.EncodeToString(pb.Bytes())
	sigBody := fmt.Sprintf(sigTpl, pubStr, sigStr)
	for i, v := range values {
		valsEnc[i] = `"` + v.String() + `"`
	}
	vals = strings.Join(valsEnc, `,`)
	return fmt.Sprintf(setArrayValuesTxTpl, msgName, from, index.String(), to.String(), slot.String(), vals, feesAmount, gasWanted, sigBody, "")
}
