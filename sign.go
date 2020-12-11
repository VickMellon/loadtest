package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/secp256k1"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

const (
	StdSignDocTpl = `{"account_number":"%d","chain_id":"%s","fee":{"amount":[{"amount":"100","denom":"uatolo"}],"gas":"200000"},"memo":"%s","msgs":[{"type":"cosmos-sdk/MsgSend","value":{"amount":[{"amount":"%s","denom":"uatolo"}],"from_address":"%s","to_address":"%s"}}],"sequence":"%d"}`
)

// StdSignDoc is replay-prevention structure.
// It includes the result of msg.GetSignBytes(),
// as well as the ChainID (prevent cross chain replay)
// and the Sequence numbers for each signature (prevent
// inchain replay and enforce tx ordering per account).
type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number" yaml:"account_number"`
	ChainID       string            `json:"chain_id" yaml:"chain_id"`
	Fee           json.RawMessage   `json:"fee" yaml:"fee"`
	Memo          string            `json:"memo" yaml:"memo"`
	Msgs          []json.RawMessage `json:"msgs" yaml:"msgs"`
	Sequence      uint64            `json:"sequence" yaml:"sequence"`
}

func signTx(tx []byte, from string, accountNumber, sequence uint64, password string) []byte {
	t1 := time.Now()
	tmpfile, err := ioutil.TempFile("", "friday_tx")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up
	if _, err := tmpfile.Write(tx); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}

	//clif tx sign tx.json --from=friday1lf9x6j0wexrs9p0spktffqdfvmkcww9jel9ygq --offline -a 0 -s 38 --indent=false
	c := exec.Command(cliFile, `tx`, `sign`, tmpfile.Name(),
		`--from=`+from, `--offline`, `--indent=false`,
		fmt.Sprintf(`--account-number=%d`, accountNumber),
		fmt.Sprintf(`--sequence=%d`, sequence),
	)
	// prepare input
	var inb bytes.Buffer
	inb.WriteString(password + "\n") // input password
	c.Stdin = &inb
	out, err := c.CombinedOutput()
	if err != nil {
		panic(err.Error() + string(out))
	}
	fmt.Println(time.Now().Sub(t1))
	return out
}

func signTxWithPk(pk crypto.PrivKey, chainId string, accountNumber, sequence uint64) []byte {
	t1 := time.Now()
	sigDoc := fmt.Sprintf(StdSignDocTpl,
		accountNumber, chainId, "", "100", "friday1pl8g6zamy8566ktzgdqtc28e06dhrh7cna783h", "friday1ag8kmam7amdv6l7xdw6w99zgw3a5a3etpvhxtr", sequence)
	sigb, err := pk.Sign([]byte(sigDoc))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(time.Now().Sub(t1))
	sig, err := json.Marshal(sigb)
	if err != nil {
		log.Fatal(err)
	}
	pb, ok := pk.PubKey().(secp256k1.PubKeySecp256k1)
	if !ok {
		log.Fatal("not secp256k1.PubKeySecp256k1")
	}
	pub, err := json.Marshal(pb[:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(pub))
	fmt.Println(string(secp256k1.PubKeyAminoName))

	return sig
}
