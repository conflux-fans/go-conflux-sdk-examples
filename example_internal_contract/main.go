package main

import (
	sdk "github.com/Conflux-Chain/go-conflux-sdk"
	"github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/conflux-fans/go-conflux-sdk-examples/context"
	exampletypes "github.com/conflux-fans/go-conflux-sdk-examples/context/types"
)

var (
	client         *sdk.Client
	config         *exampletypes.Config
	accountManger  sdk.AccountManagerOperator
	defaultAccount *types.Address
)

func init() {
	config = context.PrepareForClientExample()
	client = config.GetClient()
	accountManger = config.GetAccountManager()
	defaultAccount, _ = accountManger.GetDefault()
}

func main() {
	testAdmin()
	testStaking()
	testSponsor()
}
