package main

import (
	"fmt"
	"math/big"

	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"
	"github.com/Conflux-Chain/go-conflux-sdk/middleware"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/conflux-fans/go-conflux-sdk-examples/context"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	client := initClient()
	bulkCall(client)
	// bulkSend(client)
}

func initClient() *sdk.Client {
	config := context.PrepareForContractExample()
	sigClient := config.GetClient()
	sigClient.UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
	sigClient.UseBatchCallRpcMiddleware(middleware.BatchCallRpcConsoleMiddleware)
	return sigClient
}

func bulkCall(sigClient *sdk.Client) {

	bulkCaller := bulk.NewBulkerCaller(sigClient)

	// contract call
	mTokenBulkCaller, err := NewMyTokenBulkCaller(cfxaddress.MustNew("cfxtest:acevhhrk26e6hheja60dbr41vfmgt0t5easp2c7z46"), sigClient)
	if err != nil {
		panic(err)
	}

	addresses := []cfxaddress.Address{
		cfxaddress.MustNew("cfxtest:aamjxdgz4m84hjvf2s9rmw5uzd4dkh8aa6krdsh0ep"),
		cfxaddress.MustNew("cfxtest:aak2rra2njvd77ezwjvx04kkds9fzagfe6d5r8e957"),
		cfxaddress.MustNew("cfxtest:aamufnszywuy3bsfe1vuzbjpnd7bmpw9vutzmayap1"),
	}

	gasPrice := bulkCaller.Cfx().GetGasPrice()
	nonce0 := bulkCaller.Cfx().GetNextNonce(addresses[0])
	nonce1 := bulkCaller.Cfx().GetNextNonce(addresses[1])
	balance0 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[0].MustGetCommonAddress())
	balance1 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[1].MustGetCommonAddress())
	balance2 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[2].MustGetCommonAddress())
	name := mTokenBulkCaller.Name(*bulkCaller, nil)
	symbol := mTokenBulkCaller.Symbol(*bulkCaller, nil)
	decimals := mTokenBulkCaller.Decimals(*bulkCaller, nil)
	_struct := mTokenBulkCaller.ReturnTuple(*bulkCaller, nil)
	tupleWithStruct := mTokenBulkCaller.ReturnTupleWithStruct(*bulkCaller, nil)

	errors, err := bulkCaller.Excute()
	if err != nil {
		panic(err)
	}
	results := []interface{}{
		gasPrice, nonce0, nonce1,
		*balance0, *balance1, *balance2, *name, *symbol, *decimals,
		*_struct, *tupleWithStruct,
	}

	for i, e := range errors {
		if e != nil {
			fmt.Printf("call %vth error %v\n", i, e)
			continue
		}
		fmt.Printf("call %vth result %v\n", i, results[i])
	}
}

func bulkSend(sigClient *sdk.Client) {
	bulkSender := bulk.NewBuckSender(*sigClient)
	froms := []cfxaddress.Address{
		cfxaddress.MustNew("cfxtest:aap9kthvctunvf030rbkk9k7zbzyz12dajp1u3sp4g"),
	}
	tos := []cfxaddress.Address{
		cfxaddress.MustNew("cfxtest:aakkvauu2breh481pz49muu89bfvedz9uan0r1wp73"),
		cfxaddress.MustNew("cfxtest:aarh3pe02hdypm6pp8u3ze0shpt9vpfrrebfr1khm1"),
		cfxaddress.MustNew("cfxtest:aat6frues3g27me8p9218kgz5z2fmxzgrj7fhwgz15"),
		cfxaddress.MustNew("cfxtest:aajgaxr9wax7fuw17wae1xn7bh9w4z4496tvkzng6d"),
	}

	// cfx transfer
	bulkSender.
		AppendTransaction(newTx(&froms[0], &tos[0], types.NewBigInt(100))).
		AppendTransaction(newTx(&froms[0], &tos[1], types.NewBigInt(200))).
		AppendTransaction(newTx(&froms[0], &tos[2], types.NewBigInt(300), types.NewBigInt(2000))).
		AppendTransaction(newTx(&froms[0], &tos[3], types.NewBigInt(400))).
		AppendTransaction(newTx(&froms[0], &tos[0], types.NewBigInt(500))).
		AppendTransaction(newTx(nil, &tos[0], types.NewBigInt(500), types.NewBigInt(1))).
		AppendTransaction(newTx(nil, &tos[1], types.NewBigInt(600))).
		AppendTransaction(newTx(nil, &tos[2], types.NewBigInt(700), types.NewBigInt(1000000)))

	mTokenBulkSender, err := NewMyTokenBulkTransactor(cfxaddress.MustNew("cfxtest:acd7apn6pnfhna7w1pa8evzhwhv3085vjjp1b8bav5"), sigClient)
	if err != nil {
		panic(err)
	}

	bulkSender.
		AppendTransaction(mTokenBulkSender.Transfer(nil, tos[1].MustGetCommonAddress(), big.NewInt(10))).
		AppendTransaction(mTokenBulkSender.Transfer(nil, tos[2].MustGetCommonAddress(), big.NewInt(20))).
		AppendTransaction(mTokenBulkSender.Transfer(nil, tos[3].MustGetCommonAddress(), big.NewInt(30)))

	if err := bulkSender.PopulateTransactions(); err != nil {
		panic(err)
	}

	hashes, errors, err := bulkSender.SignAndSend()
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(hashes); i++ {
		if errors[i] != nil {
			fmt.Printf("sign and send the %vth tx error %v\n", i, errors[i])
		} else {
			fmt.Printf("the %vth tx hash %v\n", i, hashes[i])
		}
	}
}

func newTx(from *cfxaddress.Address, to *cfxaddress.Address, value *hexutil.Big, nonce ...*hexutil.Big) types.UnsignedTransaction {
	utx := types.UnsignedTransaction{}
	utx.From = from
	utx.To = to
	utx.Value = value
	if len(nonce) > 0 {
		utx.Nonce = nonce[0]
	}
	return utx
}
