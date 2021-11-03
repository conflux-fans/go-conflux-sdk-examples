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

var (
	mytokenAddr cfxaddress.Address = cfxaddress.MustNew("cfxtest:acbzmpbpb1w0a42gyzcc1gsz2wdwts2sy6fe4j208c")
)

func main() {
	client := initClient()
	bulkCall(client)
	bulkSend(client)
}

func initClient() *sdk.Client {
	config := context.PrepareForContractExample()
	sigClient := config.GetClient()
	sigClient.UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
	sigClient.UseBatchCallRpcMiddleware(middleware.BatchCallRpcConsoleMiddleware)
	return sigClient
}

func bulkCall(sigClient *sdk.Client) {

	fmt.Println("==== start test buck call")
	bulkCaller := bulk.NewBulkerCaller(sigClient)

	_defaultAcc, err := sigClient.GetAccountManager().GetDefault()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	// contract call
	mTokenBulkCaller, err := NewMyTokenBulkCaller(mytokenAddr, sigClient)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	addresses := []cfxaddress.Address{
		cfxaddress.MustNew("cfxtest:aamjxdgz4m84hjvf2s9rmw5uzd4dkh8aa6krdsh0ep"),
		cfxaddress.MustNew("cfxtest:aak2rra2njvd77ezwjvx04kkds9fzagfe6d5r8e957"),
		cfxaddress.MustNew("cfxtest:aamufnszywuy3bsfe1vuzbjpnd7bmpw9vutzmayap1"),
		*_defaultAcc,
	}

	gasPrice := bulkCaller.Cfx().GetGasPrice()
	nonce0 := bulkCaller.Cfx().GetNextNonce(addresses[0])
	nonce1 := bulkCaller.Cfx().GetNextNonce(addresses[1])
	balance0 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[0].MustGetCommonAddress())
	balance1 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[1].MustGetCommonAddress())
	balance2 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[2].MustGetCommonAddress())
	balance3 := mTokenBulkCaller.BalanceOf(*bulkCaller, nil, addresses[3].MustGetCommonAddress())
	name := mTokenBulkCaller.Name(*bulkCaller, nil)
	symbol := mTokenBulkCaller.Symbol(*bulkCaller, nil)
	decimals := mTokenBulkCaller.Decimals(*bulkCaller, nil)
	_struct := mTokenBulkCaller.ReturnTuple(*bulkCaller, nil)
	tupleWithStruct := mTokenBulkCaller.ReturnTupleWithStruct(*bulkCaller, nil)

	errors, err := bulkCaller.Excute()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}
	results := []interface{}{
		gasPrice, nonce0, nonce1,
		*balance0, *balance1, *balance2, *balance3, *name, *symbol, *decimals,
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
	fmt.Println("==== test buck send")

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

	mTokenBulkSender, err := NewMyTokenBulkTransactor(mytokenAddr, sigClient)
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	bulkSender.
		AppendTransaction(mTokenBulkSender.Transfer(nil, tos[1].MustGetCommonAddress(), big.NewInt(1))).
		AppendTransaction(mTokenBulkSender.Transfer(nil, tos[2].MustGetCommonAddress(), big.NewInt(2))).
		AppendTransaction(mTokenBulkSender.Transfer(nil, tos[3].MustGetCommonAddress(), big.NewInt(3)))

	if err := bulkSender.PopulateTransactions(); err != nil {
		panic(fmt.Sprintf("%+v", err))
	}

	hashes, errors, err := bulkSender.SignAndSend()
	if err != nil {
		panic(fmt.Sprintf("%+v", err))
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
