package main

import (
	"fmt"
	"time"
)

type localKey struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Address  string `json:"address"`
	Pubkey   string `json:"pubkey"`
	Mnemonic string `json:"mnemonic"`
}

var (
	keys map[string]localKey // indexed by name
)

func main() {
	acc := queryAccount("friday1gx8epcnd8dt64hukgmuwh4cctwuxmhrrfdgtex")
	fmt.Println(acc)
	t1 := time.Now()
	s := loadState()
	fmt.Println(time.Now().Sub(t1))
	//s.newWallet("loadtest1")
	//s.newWallet("loadtest2")
	fmt.Println(s.wallets[0].address)
	fmt.Println(s.wallets[1].address)
	//saveState(s)
	sig := signTxWithPk(s.wallets[0].privKey, "f2test", 11, 0)
	fmt.Println(string(sig))

	fmt.Println(time.Now().Sub(t1))
	return
	//
	//	pkExp := exportPrivKey("loadtest1")
	//
	//	fmt.Println(base64.StdEncoding.EncodeToString([]byte(pkExp)))
	//
	//	pk, _ := unarmorDecryptPrivKey(pkExp, "loadtest1")
	//	pb, _ := pk.PubKey().(secp256k1.PubKeySecp256k1)
	//	fmt.Println(addressFromPubKey(pb))
	//	return
	//
	//	keys = getKeysList()
	//	//fmt.Println(keys)
	//	//keys["loadtest2"] = addNewKey("loadtest2")
	//	//fmt.Println(keys)
	//	pk1, _ := unarmorDecryptPrivKey(`-----BEGIN TENDERMINT PRIVATE KEY-----
	//kdf: bcrypt
	//salt: 66C7730830513BA6BF0E318E3344E2C8
	//
	//Kwn72cMHop2R7leTdARMhjKbSq38J8PUfQtMXoJYdkezR2+nvqYXLQ6ltuJPE6FM
	//kdoLnK3xlJSZ5f9K5Ae45f7OG/cMVMBSIjk6J+I=
	//=8CTN
	//-----END TENDERMINT PRIVATE KEY-----`, "QDxseR1l")
	//
	//pk2 := privKeyFromMnemonic(`glass height scan canvas truck undo shaft core lamp fatigue toilet lemon gift phone kitten aim fantasy siege beach lens unfair worth door badge`)
	//
	//fmt.Println(pk1.Equals(pk2))
	//
	//s := addressFromPubKey(pk2)
	//fmt.Println(s)

	//acc := queryAccount("friday1pl8g6zamy8566ktzgdqtc28e06dhrh7cna783h")
	//fmt.Println(acc)
	//
	//tx := compileSendTx(keys["vick"].Address, keys["loadtest1"].Address, 100, "")
	//
	////fmt.Println(string(tx))
	//signedTx := signTx(tx, keys["vick"].Address, 9, 2, "QDxseR1l")
	//fmt.Println(string(signedTx))
	////zeoValNTduFXl+rzwK3X9kxDh235E/v7GRlhPsKZTHhaohoftvNdGfV7xR/0Lqu9k2MFrsp+5UoBe7hPxwppCQ==
	//
	//signedTx2 := signTxWithPk(pk1, "f2test", "9", "2")
	//fmt.Println(string(signedTx2))
	////nGqjnuJ9oW8Z2U7t2Ev5Lcxg6h+uGn2j9k+NBJ0bqvkpl850u86ywosRcVSFRh4dibtgJLfmAD07JAO7mjSPcg==

}

func compileSendTx(from, to string, amount uint64, memo string) []byte {
	return []byte(fmt.Sprintf(sendTxTpl, from, to, amount, memo))
}
