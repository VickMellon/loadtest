package main

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
)

const (
	AccountAddressPrefix = `eth`
)

func newPrivKey() *ethsecp256k1.PrivKey {
	pk, _ := ethsecp256k1.GenerateKey()
	return pk
}

func addressFromPubKey(pb cryptotypes.PubKey) (string, error) {
	addr, err := bech32.ConvertAndEncode(AccountAddressPrefix, pb.Address())
	if err != nil {
		return "", err
	}
	return addr, nil
}
