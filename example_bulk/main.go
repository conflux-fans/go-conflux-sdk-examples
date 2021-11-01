package main

import (
	"fmt"

	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"
	"github.com/Conflux-Chain/go-conflux-sdk/example/context"
	"github.com/Conflux-Chain/go-conflux-sdk/middleware"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/Conflux-Chain/go-conflux-sdk/types/cfxaddress"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func main() {
	config := context.PrepareForContractExample()
	sigClient := config.GetClient()
	sigClient.UseCallRpcMiddleware(middleware.CallRpcConsoleMiddleware)
	sigClient.UseBatchCallRpcMiddleware(middleware.BatchCallRpcConsoleMiddleware)

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

	bulkSender.
		AppendTransaction(newTx(&froms[0], &tos[0], types.NewBigInt(100))).
		AppendTransaction(newTx(&froms[0], &tos[1], types.NewBigInt(200))).
		AppendTransaction(newTx(&froms[0], &tos[2], types.NewBigInt(300), types.NewBigInt(2000))).
		AppendTransaction(newTx(&froms[0], &tos[3], types.NewBigInt(400))).
		AppendTransaction(newTx(&froms[0], &tos[0], types.NewBigInt(500))).
		AppendTransaction(newTx(nil, &tos[0], types.NewBigInt(500), types.NewBigInt(1))).
		AppendTransaction(newTx(nil, &tos[1], types.NewBigInt(600))).
		AppendTransaction(newTx(nil, &tos[2], types.NewBigInt(700), types.NewBigInt(1000000)))

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
