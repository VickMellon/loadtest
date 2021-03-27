package main

import (
	"github.com/cosmos/ethermint/crypto/ethsecp256k1"
	"github.com/tendermint/tendermint/libs/bech32"
)

const (
	AccountAddressPrefix = `eth`
)

func newPrivKey() ethsecp256k1.PrivKey {
	pk, _ := ethsecp256k1.GenerateKey()
	return pk
}

func addressFromPubKey(pb ethsecp256k1.PubKey) (string, error) {
	addr, err := bech32.ConvertAndEncode(AccountAddressPrefix, pb.Address())
	if err != nil {
		return "", err
	}
	return addr, nil
}
