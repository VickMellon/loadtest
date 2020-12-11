package main

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/hd"
	"github.com/cosmos/go-bip39"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"github.com/tendermint/tendermint/libs/bech32"
	"log"
	"os/exec"
	"time"
)

const (
	AccountAddressPrefix = `friday`
)

func newPrivKey() secp256k1.PrivKeySecp256k1 {
	return secp256k1.GenPrivKey()
}

func exportPrivKey(name string) string {
	t1 := time.Now()
	c := exec.Command(cliFile, `keys`, `export`, name)
	// prepare input
	var inb bytes.Buffer
	inb.WriteString(name + "\n") // input password to decrypt from storage (contract: same as name)
	inb.WriteString(name + "\n") // input password to encrypt exported (contract: same as name)
	c.Stdin = &inb
	// for some reason clif output exported key to StdErr instead of StdOut
	out, err := c.CombinedOutput()
	if err != nil {
		log.Fatal(err, out)
	}
	fmt.Println(time.Now().Sub(t1))
	return string(out)
}

func privKeyFromMnemonic(mn string) secp256k1.PrivKeySecp256k1 {
	seed, err := bip39.NewSeedWithErrorChecking(mn, "")
	if err != nil {
		log.Fatal(err)
	}
	// create master key and derive first key:
	masterPriv, ch := hd.ComputeMastersFromSeed(seed)
	hdPath := hd.NewFundraiserParams(0, 118, 0)
	derivedPriv, err := hd.DerivePrivateKeyForPath(masterPriv, ch, hdPath.String())
	if err != nil {
		log.Fatal(err)
	}
	return derivedPriv
}

func addressFromPubKey(pb secp256k1.PubKeySecp256k1) (string, error) {
	addr, err := bech32.ConvertAndEncode(AccountAddressPrefix, pb.Address())
	if err != nil {
		return "", err
	}
	return addr, nil
}
