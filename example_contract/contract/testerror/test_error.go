// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package testerror

import (
	"math/big"
	"strings"

	"github.com/Conflux-Chain/conflux-abigen/bind"
	"github.com/Conflux-Chain/go-conflux-sdk/cfxclient/bulk"

	types "github.com/Conflux-Chain/go-conflux-sdk/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethBind "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = ethBind.Bind
	_ = common.Big1
	_ = ethtypes.BloomLookup
	_ = event.NewSubscription
)

// TestErrorABI is the input ABI used to generate the binding from.
const TestErrorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"msg\",\"type\":\"string\"}],\"name\":\"Error\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mustRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mustRevertByCall\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// TestErrorFuncSigs maps the 4-byte function signature to its string representation.
var TestErrorFuncSigs = map[string]string{
	"08c379a0": "Error(string)",
	"91c5eace": "mustRevert()",
	"ba050c11": "mustRevertByCall()",
}

// TestErrorBin is the compiled bytecode used for deploying new contracts.
var TestErrorBin = "0x608060405234801561001057600080fd5b506101e3806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806308c379a01461004657806391c5eace14610059578063ba050c1114610061575b600080fd5b6100576100543660046100e6565b50565b005b610057610069565b61005761009e565b60405162461bcd60e51b815260206004820152600560248201526422b93937b960d91b60448201526064015b60405180910390fd5b60405162461bcd60e51b815260206004820152601860248201527f4572726f72207768656e207265616420636f6e747261637400000000000000006044820152606401610095565b6000602082840312156100f857600080fd5b813567ffffffffffffffff8082111561011057600080fd5b818401915084601f83011261012457600080fd5b81358181111561013657610136610197565b604051601f8201601f19908116603f0116810190838211818310171561015e5761015e610197565b8160405282815287602084870101111561017757600080fd5b826020860160208301376000928101602001929092525095945050505050565b634e487b7160e01b600052604160045260246000fdfea2646970667358221220313db9fb5606a22cdeb09628e94059c37c687e220b1a5f8860f82812046fc3c764736f6c63430008060033"

// DeployTestError deploys a new Conflux contract, binding an instance of TestError to it.
func DeployTestError(auth *bind.TransactOpts, backend bind.ContractBackend) (*types.UnsignedTransaction, *types.Hash, *TestError, error) {
	parsed, err := abi.JSON(strings.NewReader(TestErrorABI))
	if err != nil {
		return nil, nil, nil, err
	}

	tx, hash, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TestErrorBin), backend)
	if err != nil {
		return nil, nil, nil, err
	}
	return tx, hash, &TestError{TestErrorCaller: TestErrorCaller{contract: contract}, TestErrorTransactor: TestErrorTransactor{contract: contract}, TestErrorFilterer: TestErrorFilterer{contract: contract}}, nil
}

// TestError is an auto generated Go binding around an Conflux contract.
type TestError struct {
	TestErrorCaller     // Read-only binding to the contract
	TestErrorTransactor // Write-only binding to the contract
	TestErrorFilterer   // Log filterer for contract events
}

// TestErrorCaller is an auto generated read-only Go binding around an Conflux contract.
type TestErrorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorBulkCaller is an auto generated read-only Go binding around an Conflux contract.
type TestErrorBulkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorTransactor is an auto generated write-only Go binding around an Conflux contract.
type TestErrorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorBulkTransactor is an auto generated write-only Go binding around an Conflux contract.
type TestErrorBulkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorFilterer is an auto generated log filtering Go binding around an Conflux contract events.
type TestErrorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TestErrorSession is an auto generated Go binding around an Conflux contract,
// with pre-set call and transact options.
type TestErrorSession struct {
	Contract     *TestError        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TestErrorCallerSession is an auto generated read-only Go binding around an Conflux contract,
// with pre-set call options.
type TestErrorCallerSession struct {
	Contract *TestErrorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TestErrorTransactorSession is an auto generated write-only Go binding around an Conflux contract,
// with pre-set transact options.
type TestErrorTransactorSession struct {
	Contract     *TestErrorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TestErrorRaw is an auto generated low-level Go binding around an Conflux contract.
type TestErrorRaw struct {
	Contract *TestError // Generic contract binding to access the raw methods on
}

// TestErrorCallerRaw is an auto generated low-level read-only Go binding around an Conflux contract.
type TestErrorCallerRaw struct {
	Contract *TestErrorCaller // Generic read-only contract binding to access the raw methods on
}

// TestErrorTransactorRaw is an auto generated low-level write-only Go binding around an Conflux contract.
type TestErrorTransactorRaw struct {
	Contract *TestErrorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTestError creates a new instance of TestError, bound to a specific deployed contract.
func NewTestError(address types.Address, backend bind.ContractBackend) (*TestError, error) {
	contract, err := bindTestError(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TestError{TestErrorCaller: TestErrorCaller{contract: contract}, TestErrorTransactor: TestErrorTransactor{contract: contract}, TestErrorFilterer: TestErrorFilterer{contract: contract}}, nil
}

// NewTestErrorCaller creates a new read-only instance of TestError, bound to a specific deployed contract.
func NewTestErrorCaller(address types.Address, caller bind.ContractCaller) (*TestErrorCaller, error) {
	contract, err := bindTestError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestErrorCaller{contract: contract}, nil
}

// NewTestErrorTransactor creates a new write-only instance of TestError, bound to a specific deployed contract.
func NewTestErrorTransactor(address types.Address, transactor bind.ContractTransactor) (*TestErrorTransactor, error) {
	contract, err := bindTestError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestErrorTransactor{contract: contract}, nil
}

// NewTestErrorFilterer creates a new log filterer instance of TestError, bound to a specific deployed contract.
func NewTestErrorFilterer(address types.Address, filterer bind.ContractFilterer) (*TestErrorFilterer, error) {
	contract, err := bindTestError(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TestErrorFilterer{contract: contract}, nil
}

// NewTestErrorCaller creates a new read-only instance of TestError, bound to a specific deployed contract.
func NewTestErrorBulkCaller(address types.Address, caller bind.ContractCaller) (*TestErrorBulkCaller, error) {
	contract, err := bindTestError(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TestErrorBulkCaller{contract: contract}, nil
}

// NewTestErrorBulkTransactor creates a new write-only instance of TestError, bound to a specific deployed contract.
func NewTestErrorBulkTransactor(address types.Address, transactor bind.ContractTransactor) (*TestErrorBulkTransactor, error) {
	contract, err := bindTestError(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TestErrorBulkTransactor{contract: contract}, nil
}

// bindTestError binds a generic wrapper to an already deployed contract.
func bindTestError(address types.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TestErrorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestError *TestErrorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestError.Contract.TestErrorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestError *TestErrorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.Contract.TestErrorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestError *TestErrorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.Contract.TestErrorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TestError *TestErrorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TestError.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TestError *TestErrorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TestError *TestErrorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.Contract.contract.Transact(opts, method, params...)
}

// Error is a free data retrieval call binding the contract method 0x08c379a0.
//
// Solidity: function Error(string msg) pure returns()
func (_TestError *TestErrorCaller) Error(opts *bind.CallOpts, msg string) error {
	var out []interface{}
	__err := _TestError.contract.Call(opts, &out, "Error", msg)

	if __err != nil {
		return __err
	}

	return __err

}

// Error is a free data retrieval call binding the contract method 0x08c379a0.
//
// Solidity: function Error(string msg) pure returns()
func (_TestError *TestErrorBulkCaller) Error(bulkcaller bulk.BulkCaller, opts *bind.CallOpts, msg string) *error {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _TestError.contract.GenRequest(opts, "Error", msg)

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _TestError.contract.DecodeOutput(&out, rawOut, "Error")
		if err != nil {
			return err
		}

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return __err

}

// Error is a free data retrieval call binding the contract method 0x08c379a0.
//
// Solidity: function Error(string msg) pure returns()
func (_TestError *TestErrorSession) Error(msg string) error {
	return _TestError.Contract.Error(&_TestError.CallOpts, msg)
}

// Error is a free data retrieval call binding the contract method 0x08c379a0.
//
// Solidity: function Error(string msg) pure returns()
func (_TestError *TestErrorCallerSession) Error(msg string) error {
	return _TestError.Contract.Error(&_TestError.CallOpts, msg)
}

// MustRevertByCall is a free data retrieval call binding the contract method 0xba050c11.
//
// Solidity: function mustRevertByCall() pure returns()
func (_TestError *TestErrorCaller) MustRevertByCall(opts *bind.CallOpts) error {
	var out []interface{}
	__err := _TestError.contract.Call(opts, &out, "mustRevertByCall")

	if __err != nil {
		return __err
	}

	return __err

}

// MustRevertByCall is a free data retrieval call binding the contract method 0xba050c11.
//
// Solidity: function mustRevertByCall() pure returns()
func (_TestError *TestErrorBulkCaller) MustRevertByCall(bulkcaller bulk.BulkCaller, opts *bind.CallOpts) *error {

	if opts == nil {
		opts = new(bind.CallOpts)
	}
	__request := _TestError.contract.GenRequest(opts, "mustRevertByCall")

	__err := new(error)

	outDecoder := func(rawOut []byte) error {
		out := []interface{}{}
		err := _TestError.contract.DecodeOutput(&out, rawOut, "mustRevertByCall")
		if err != nil {
			return err
		}

		return nil
	}

	bulkcaller.Customer().ContractCall(__request, opts.EpochNumber, outDecoder, __err)

	return __err

}

// MustRevertByCall is a free data retrieval call binding the contract method 0xba050c11.
//
// Solidity: function mustRevertByCall() pure returns()
func (_TestError *TestErrorSession) MustRevertByCall() error {
	return _TestError.Contract.MustRevertByCall(&_TestError.CallOpts)
}

// MustRevertByCall is a free data retrieval call binding the contract method 0xba050c11.
//
// Solidity: function mustRevertByCall() pure returns()
func (_TestError *TestErrorCallerSession) MustRevertByCall() error {
	return _TestError.Contract.MustRevertByCall(&_TestError.CallOpts)
}

// MustRevert is a paid mutator transaction binding the contract method 0x91c5eace.
//
// Solidity: function mustRevert() returns()
func (_TestError *TestErrorTransactor) MustRevert(opts *bind.TransactOpts) (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.contract.Transact(opts, "mustRevert")
}

// MustRevert is a paid mutator transaction binding the contract method 0x91c5eace.
//
// Solidity: function mustRevert() returns()
func (_TestError *TestErrorBulkTransactor) MustRevert(opts *bind.TransactOpts) types.UnsignedTransaction {
	return _TestError.contract.GenUnsignedTransaction(opts, "mustRevert")
}

// MustRevert is a paid mutator transaction binding the contract method 0x91c5eace.
//
// Solidity: function mustRevert() returns()
func (_TestError *TestErrorSession) MustRevert() (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.Contract.MustRevert(&_TestError.TransactOpts)
}

// MustRevert is a paid mutator transaction binding the contract method 0x91c5eace.
//
// Solidity: function mustRevert() returns()
func (_TestError *TestErrorTransactorSession) MustRevert() (*types.UnsignedTransaction, *types.Hash, error) {
	return _TestError.Contract.MustRevert(&_TestError.TransactOpts)
}
