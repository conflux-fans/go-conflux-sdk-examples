package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"time"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
)

var am *sdk.AccountManager

func init() {
	initAccountManager()
	fmt.Println("init account manager done")
}

func main() {
	creatAccount()
	importAccount()
	listAccount()
	signTx()
	decodeRawTx()
	updateAccount()
	deleteAccount()
}

func initAccountManager() *sdk.AccountManager {
	keydir := "./keystore"
	am = sdk.NewAccountManager(keydir, 1)
	return am
}

func listAccount() {
	fmt.Printf("account list: %+v\n\n", am.List())
}

func creatAccount() {
	am := initAccountManager()
	addr, err := am.Create("hello")
	if err != nil {
		fmt.Println("create account error", err)
		return
	}
	fmt.Println("creat account:", addr)
}

func importAccount() {
	am := initAccountManager()
	dir, _ := os.Getwd()

	filePath := path.Join(dir, "keystore", fmt.Sprint(time.Now().Unix()))
	content := `{"address":"16fd16cb6dbcba8ad8e91221e57f3be9160282a9","crypto":{"cipher":"aes-128-ctr","ciphertext":"2deb8e03403638b6b13edfc6cc0e925821f7e576202cfb9cd984f2d85a9c05d7","cipherparams":{"iv":"05db211d3302f5c64105bffbb8d66f61"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"df7f2bbab8e980dfe15a82e1cfec567b0f2f240baa002b7e78ca7f34a7703c68"},"mac":"9e8a705491a52052473a53de5286c87ed27fda0f6ec2813c29875c4d0f7b4814"},"id":"3df031eb-960f-4e32-9bea-fdb8289d457a","version":3}`
	if err := ioutil.WriteFile(filePath, []byte(content), 0777); err != nil {
		panic(err)
	}

	addr, err := am.Import(filePath, "hello", "hello")
	if err != nil {
		fmt.Println("import account error:", err)
		return
	}
	fmt.Println("import account done:", addr)
}

func updateAccount() {
	address := cfxaddress.MustNewFromHex("0x16fd16cb6dbcba8ad8e91221e57f3be9160282a9", 1)
	err := am.Update(address, "hello", "hello world")
	if err != nil {
		fmt.Printf("update address error: %v \n\n", err)
		return
	}
	fmt.Printf("update address %v done\n\n", address)
}

func deleteAccount() {
	address := cfxaddress.MustNewFromHex("0x16fd16cb6dbcba8ad8e91221e57f3be9160282a9", 1)
	err := am.Delete(address, "hello world")
	if err != nil {
		fmt.Printf("delete address error: %v \n\n", err)
		return
	}
	fmt.Printf("delete address %v done\n\n", address)
}

func signTx() []byte {
	am := initAccountManager()

	from := cfxaddress.MustNewFromHex("0x16fd16cb6dbcba8ad8e91221e57f3be9160282a9", 1)
	to := cfxaddress.MustNewFromHex("0x10f4bcf113e0b896d9b34294fd3da86b4adf0302", 1)
	unSignedTx := types.UnsignedTransaction{
		UnsignedTransactionBase: types.UnsignedTransactionBase{
			From:        &from,
			Value:       types.NewBigInt(100),
			Gas:         types.NewBigInt(21000),
			GasPrice:    types.NewBigInt(100000000),
			Nonce:       types.NewBigInt(1),
			EpochHeight: types.NewUint64(1),
			ChainID:     types.NewUint(1),
		},
		To: &to,
	}

	signedTx, err := am.SignAndEcodeTransactionWithPassphrase(unSignedTx, "hello")
	if err != nil {
		fmt.Printf("signed tx %+v error:%v\n\n", unSignedTx, err)
		return nil
	}
	fmt.Printf("signed tx %+v result:\n0x%x\n\n", unSignedTx, signedTx)
	return signedTx
}

func decodeRawTx() {
	rawTx, _ := hex.DecodeString("f867e3018405f5e1008252089410f4bcf113e0b896d9b34294fd3da86b4adf0302648001018080a072aa2777c4b8cee3829ea3fb9658276e40cc4234eb05f176459042e48f69428da07a9bbee20b9a219907c91b562b64ee2e9456d2f280c31ce98736d0feb5e47372")
	tx := new(types.SignedTransaction)
	err := tx.Decode(rawTx, 1)
	if err != nil {
		fmt.Printf("decoded rawTx error: %+v\n\n", err)
		return
	}
	fmt.Printf("decoded rawTx %x result: %+v\n\n", rawTx, tx)
}
