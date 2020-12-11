package main

import "github.com/tendermint/tendermint/crypto"

const (
	sendTxTpl      = `{"type":"cosmos-sdk/StdTx","value":{"msg":[{"type":"cosmos-sdk/MsgSend","value":{"from_address":"%s","to_address":"%s","amount":[{"denom":"uatolo","amount":"%d"}]}}],"fee":{"amount":[{"denom":"uatolo","amount":"100"}],"gas":"200000"},"signatures":null,"memo":"%s"}}`
	broadcastTxTpl = `{"mode":"%s","tx":"%s"}`
)

type Transfer struct {
	From   string `json:"from"`   // bech32 accAddress
	To     string `json:"to"`     // bech32 accAddress
	Amount string `json:"amount"` // sum in uatolo
}

type StdTx struct {
	Msgs       []MsgSend      `json:"msg"`
	Fee        StdFee         `json:"fee"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
}
type MsgSend struct {
	Amount      Coins  `json:"amount"`
	FromAddress string `json:"from_address"`
	ToAddress   string `json:"to_address"`
}
type StdSignature struct {
	crypto.PubKey `json:"pub_key"` // optional
	Signature     []byte           `json:"signature"`
}
type StdFee struct {
	Amount Coins  `json:"amount"`
	Gas    string `json:"gas"`
}
type Coins []Coin
type Coin struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}
