package main

import (
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/libs/bech32"
)

const (
	AccountAddressPrefix = `friday`
)

func newPrivKey() secp256k1.PrivKeySecp256k1 {
	return secp256k1.GenPrivKey()
}

func addressFromPubKey(pb secp256k1.PubKeySecp256k1) (string, error) {
	addr, err := bech32.ConvertAndEncode(AccountAddressPrefix, pb.Address())
	if err != nil {
		return "", err
	}
	return addr, nil
}
