// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IndexerABI is the input ABI used to generate the binding from.
const IndexerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"Add\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"form\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AlterResolver\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"form\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"}],\"name\":\"AlterOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"data\",\"type\":\"string\"}],\"name\":\"Error\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"add\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"}],\"name\":\"getResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"alterResolver\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"key\",\"type\":\"string\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"alterOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IndexerBin is the compiled bytecode used for deploying new contracts.
const IndexerBin = `0x608060405234801561001057600080fd5b506113f5806100206000396000f3fe608060405260043610610078576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632bffc7ed1461007d5780634aaf4a121461017d578063693ec85e14610285578063cbdc3fe1146103c0578063e4ae7d77146104c0578063f60b53e2146105c8575b600080fd5b34801561008957600080fd5b50610163600480360360408110156100a057600080fd5b81019080803590602001906401000000008111156100bd57600080fd5b8201836020820111156100cf57600080fd5b803590602001918460018302840111640100000000831117156100f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506106c8565b604051808215151515815260200191505060405180910390f35b34801561018957600080fd5b50610243600480360360208110156101a057600080fd5b81019080803590602001906401000000008111156101bd57600080fd5b8201836020820111156101cf57600080fd5b803590602001918460018302840111640100000000831117156101f157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610a67565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561029157600080fd5b5061034b600480360360208110156102a857600080fd5b81019080803590602001906401000000008111156102c557600080fd5b8201836020820111156102d757600080fd5b803590602001918460018302840111640100000000831117156102f957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610afe565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019250505060405180910390f35b3480156103cc57600080fd5b506104a6600480360360408110156103e357600080fd5b810190808035906020019064010000000081111561040057600080fd5b82018360208201111561041257600080fd5b8035906020019184600183028401116401000000008311171561043457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c28565b604051808215151515815260200191505060405180910390f35b3480156104cc57600080fd5b50610586600480360360208110156104e357600080fd5b810190808035906020019064010000000081111561050057600080fd5b82018360208201111561051257600080fd5b8035906020019184600183028401116401000000008311171561053457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610fad565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156105d457600080fd5b506106ae600480360360408110156105eb57600080fd5b810190808035906020019064010000000081111561060857600080fd5b82018360208201111561061a57600080fd5b8035906020019184600183028401116401000000008311171561063c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611044565b604051808215151515815260200191505060405180910390f35b60008073ffffffffffffffffffffffffffffffffffffffff166000846040518082805190602001908083835b60208310151561071957805182526020820191506020810190506020830392506106f4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156107fe577f08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa6040518080602001828103825260208152602001807fe5b7b2e69c89e6ada4e5908de7a7b0e5afb9e5ba94e79a847265736f6c76657281525060200191505060405180910390a160009050610a61565b336000846040518082805190602001908083835b6020831015156108375780518252602082019150602081019050602083039250610812565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816000846040518082805190602001908083835b6020831015156108e457805182526020820191506020810190506020830392506108bf565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fec689a3871c35587e4800f14216f987ee744b924aff21741edc2e167e2dd43e883338460405180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b83811015610a20578082015181840152602081019050610a05565b50505050905090810190601f168015610a4d5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a1600190505b92915050565b600080826040518082805190602001908083835b602083101515610aa05780518252602082019150602081019050602083039250610a7b565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b6000806000836040518082805190602001908083835b602083101515610b395780518252602082019150602081019050602083039250610b14565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000846040518082805190602001908083835b602083101515610bc85780518252602082019150602081019050602083039250610ba3565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691509150915091565b60003373ffffffffffffffffffffffffffffffffffffffff166000846040518082805190602001908083835b602083101515610c795780518252602082019150602081019050602083039250610c54565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610d5e577f08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa6040518080602001828103825260188152602001807fe4bda0e4b88de698afe6ada46b6579e79a84e4b8bbe4baba000000000000000081525060200191505060405180910390a160009050610fa7565b600080846040518082805190602001908083835b602083101515610d975780518252602082019150602081019050602083039250610d72565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050826000856040518082805190602001908083835b602083101515610e295780518252602082019150602081019050602083039250610e04565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f0a7047ba8be4d874e67aebc953a70ff6db03a81782549290ac646e0738ddfc0484828560405180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b83811015610f65578082015181840152602081019050610f4a565b50505050905090810190601f168015610f925780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160019150505b92915050565b600080826040518082805190602001908083835b602083101515610fe65780518252602082019150602081019050602083039250610fc1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60003373ffffffffffffffffffffffffffffffffffffffff166000846040518082805190602001908083835b6020831015156110955780518252602082019150602081019050602083039250611070565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561117a577f08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa6040518080602001828103825260188152602001807fe4bda0e4b88de698afe6ada46b6579e79a84e4b8bbe4baba000000000000000081525060200191505060405180910390a1600090506113c3565b600080846040518082805190602001908083835b6020831015156111b3578051825260208201915060208101905060208303925061118e565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050826000856040518082805190602001908083835b6020831015156112455780518252602082019150602081019050602083039250611220565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f46bd035a76a8302bb74520f9226b59925d8186784298f88ad636a4ea46b85b2184828560405180806020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828103825285818151815260200191508051906020019080838360005b83811015611381578082015181840152602081019050611366565b50505050905090810190601f1680156113ae5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a160019150505b9291505056fea165627a7a72305820eceecec8f082306a9914267a9aa8c1c321e8f776f7d09691122198d42017025d0029`

// DeployIndexer deploys a new Ethereum contract, binding an instance of Indexer to it.
func DeployIndexer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Indexer, error) {
	parsed, err := abi.JSON(strings.NewReader(IndexerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IndexerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Indexer{IndexerCaller: IndexerCaller{contract: contract}, IndexerTransactor: IndexerTransactor{contract: contract}, IndexerFilterer: IndexerFilterer{contract: contract}}, nil
}

// Indexer is an auto generated Go binding around an Ethereum contract.
type Indexer struct {
	IndexerCaller     // Read-only binding to the contract
	IndexerTransactor // Write-only binding to the contract
	IndexerFilterer   // Log filterer for contract events
}

// IndexerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IndexerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IndexerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IndexerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IndexerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IndexerSession struct {
	Contract     *Indexer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IndexerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IndexerCallerSession struct {
	Contract *IndexerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IndexerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IndexerTransactorSession struct {
	Contract     *IndexerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IndexerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IndexerRaw struct {
	Contract *Indexer // Generic contract binding to access the raw methods on
}

// IndexerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IndexerCallerRaw struct {
	Contract *IndexerCaller // Generic read-only contract binding to access the raw methods on
}

// IndexerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IndexerTransactorRaw struct {
	Contract *IndexerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIndexer creates a new instance of Indexer, bound to a specific deployed contract.
func NewIndexer(address common.Address, backend bind.ContractBackend) (*Indexer, error) {
	contract, err := bindIndexer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Indexer{IndexerCaller: IndexerCaller{contract: contract}, IndexerTransactor: IndexerTransactor{contract: contract}, IndexerFilterer: IndexerFilterer{contract: contract}}, nil
}

// NewIndexerCaller creates a new read-only instance of Indexer, bound to a specific deployed contract.
func NewIndexerCaller(address common.Address, caller bind.ContractCaller) (*IndexerCaller, error) {
	contract, err := bindIndexer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IndexerCaller{contract: contract}, nil
}

// NewIndexerTransactor creates a new write-only instance of Indexer, bound to a specific deployed contract.
func NewIndexerTransactor(address common.Address, transactor bind.ContractTransactor) (*IndexerTransactor, error) {
	contract, err := bindIndexer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IndexerTransactor{contract: contract}, nil
}

// NewIndexerFilterer creates a new log filterer instance of Indexer, bound to a specific deployed contract.
func NewIndexerFilterer(address common.Address, filterer bind.ContractFilterer) (*IndexerFilterer, error) {
	contract, err := bindIndexer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IndexerFilterer{contract: contract}, nil
}

// bindIndexer binds a generic wrapper to an already deployed contract.
func bindIndexer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IndexerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Indexer *IndexerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Indexer.Contract.IndexerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Indexer *IndexerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Indexer.Contract.IndexerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Indexer *IndexerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Indexer.Contract.IndexerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Indexer *IndexerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Indexer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Indexer *IndexerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Indexer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Indexer *IndexerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Indexer.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(address, address)
func (_Indexer *IndexerCaller) Get(opts *bind.CallOpts, key string) (common.Address, common.Address, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(common.Address)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Indexer.contract.Call(opts, out, "get", key)
	return *ret0, *ret1, err
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(address, address)
func (_Indexer *IndexerSession) Get(key string) (common.Address, common.Address, error) {
	return _Indexer.Contract.Get(&_Indexer.CallOpts, key)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string key) constant returns(address, address)
func (_Indexer *IndexerCallerSession) Get(key string) (common.Address, common.Address, error) {
	return _Indexer.Contract.Get(&_Indexer.CallOpts, key)
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string key) constant returns(address)
func (_Indexer *IndexerCaller) GetOwner(opts *bind.CallOpts, key string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Indexer.contract.Call(opts, out, "getOwner", key)
	return *ret0, err
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string key) constant returns(address)
func (_Indexer *IndexerSession) GetOwner(key string) (common.Address, error) {
	return _Indexer.Contract.GetOwner(&_Indexer.CallOpts, key)
}

// GetOwner is a free data retrieval call binding the contract method 0x4aaf4a12.
//
// Solidity: function getOwner(string key) constant returns(address)
func (_Indexer *IndexerCallerSession) GetOwner(key string) (common.Address, error) {
	return _Indexer.Contract.GetOwner(&_Indexer.CallOpts, key)
}

// GetResolver is a free data retrieval call binding the contract method 0xe4ae7d77.
//
// Solidity: function getResolver(string key) constant returns(address)
func (_Indexer *IndexerCaller) GetResolver(opts *bind.CallOpts, key string) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Indexer.contract.Call(opts, out, "getResolver", key)
	return *ret0, err
}

// GetResolver is a free data retrieval call binding the contract method 0xe4ae7d77.
//
// Solidity: function getResolver(string key) constant returns(address)
func (_Indexer *IndexerSession) GetResolver(key string) (common.Address, error) {
	return _Indexer.Contract.GetResolver(&_Indexer.CallOpts, key)
}

// GetResolver is a free data retrieval call binding the contract method 0xe4ae7d77.
//
// Solidity: function getResolver(string key) constant returns(address)
func (_Indexer *IndexerCallerSession) GetResolver(key string) (common.Address, error) {
	return _Indexer.Contract.GetResolver(&_Indexer.CallOpts, key)
}

// Add is a paid mutator transaction binding the contract method 0x2bffc7ed.
//
// Solidity: function add(string key, address resolver) returns(bool)
func (_Indexer *IndexerTransactor) Add(opts *bind.TransactOpts, key string, resolver common.Address) (*types.Transaction, error) {
	return _Indexer.contract.Transact(opts, "add", key, resolver)
}

// Add is a paid mutator transaction binding the contract method 0x2bffc7ed.
//
// Solidity: function add(string key, address resolver) returns(bool)
func (_Indexer *IndexerSession) Add(key string, resolver common.Address) (*types.Transaction, error) {
	return _Indexer.Contract.Add(&_Indexer.TransactOpts, key, resolver)
}

// Add is a paid mutator transaction binding the contract method 0x2bffc7ed.
//
// Solidity: function add(string key, address resolver) returns(bool)
func (_Indexer *IndexerTransactorSession) Add(key string, resolver common.Address) (*types.Transaction, error) {
	return _Indexer.Contract.Add(&_Indexer.TransactOpts, key, resolver)
}

// AlterOwner is a paid mutator transaction binding the contract method 0xf60b53e2.
//
// Solidity: function alterOwner(string key, address owner) returns(bool)
func (_Indexer *IndexerTransactor) AlterOwner(opts *bind.TransactOpts, key string, owner common.Address) (*types.Transaction, error) {
	return _Indexer.contract.Transact(opts, "alterOwner", key, owner)
}

// AlterOwner is a paid mutator transaction binding the contract method 0xf60b53e2.
//
// Solidity: function alterOwner(string key, address owner) returns(bool)
func (_Indexer *IndexerSession) AlterOwner(key string, owner common.Address) (*types.Transaction, error) {
	return _Indexer.Contract.AlterOwner(&_Indexer.TransactOpts, key, owner)
}

// AlterOwner is a paid mutator transaction binding the contract method 0xf60b53e2.
//
// Solidity: function alterOwner(string key, address owner) returns(bool)
func (_Indexer *IndexerTransactorSession) AlterOwner(key string, owner common.Address) (*types.Transaction, error) {
	return _Indexer.Contract.AlterOwner(&_Indexer.TransactOpts, key, owner)
}

// AlterResolver is a paid mutator transaction binding the contract method 0xcbdc3fe1.
//
// Solidity: function alterResolver(string key, address resolver) returns(bool)
func (_Indexer *IndexerTransactor) AlterResolver(opts *bind.TransactOpts, key string, resolver common.Address) (*types.Transaction, error) {
	return _Indexer.contract.Transact(opts, "alterResolver", key, resolver)
}

// AlterResolver is a paid mutator transaction binding the contract method 0xcbdc3fe1.
//
// Solidity: function alterResolver(string key, address resolver) returns(bool)
func (_Indexer *IndexerSession) AlterResolver(key string, resolver common.Address) (*types.Transaction, error) {
	return _Indexer.Contract.AlterResolver(&_Indexer.TransactOpts, key, resolver)
}

// AlterResolver is a paid mutator transaction binding the contract method 0xcbdc3fe1.
//
// Solidity: function alterResolver(string key, address resolver) returns(bool)
func (_Indexer *IndexerTransactorSession) AlterResolver(key string, resolver common.Address) (*types.Transaction, error) {
	return _Indexer.Contract.AlterResolver(&_Indexer.TransactOpts, key, resolver)
}

// IndexerAddIterator is returned from FilterAdd and is used to iterate over the raw logs and unpacked data for Add events raised by the Indexer contract.
type IndexerAddIterator struct {
	Event *IndexerAdd // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IndexerAddIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexerAdd)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IndexerAdd)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IndexerAddIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexerAddIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexerAdd represents a Add event raised by the Indexer contract.
type IndexerAdd struct {
	Key      string
	Owner    common.Address
	Resolver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdd is a free log retrieval operation binding the contract event 0xec689a3871c35587e4800f14216f987ee744b924aff21741edc2e167e2dd43e8.
//
// Solidity: event Add(string key, address owner, address resolver)
func (_Indexer *IndexerFilterer) FilterAdd(opts *bind.FilterOpts) (*IndexerAddIterator, error) {

	logs, sub, err := _Indexer.contract.FilterLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return &IndexerAddIterator{contract: _Indexer.contract, event: "Add", logs: logs, sub: sub}, nil
}

// WatchAdd is a free log subscription operation binding the contract event 0xec689a3871c35587e4800f14216f987ee744b924aff21741edc2e167e2dd43e8.
//
// Solidity: event Add(string key, address owner, address resolver)
func (_Indexer *IndexerFilterer) WatchAdd(opts *bind.WatchOpts, sink chan<- *IndexerAdd) (event.Subscription, error) {

	logs, sub, err := _Indexer.contract.WatchLogs(opts, "Add")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexerAdd)
				if err := _Indexer.contract.UnpackLog(event, "Add", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IndexerAlterOwnerIterator is returned from FilterAlterOwner and is used to iterate over the raw logs and unpacked data for AlterOwner events raised by the Indexer contract.
type IndexerAlterOwnerIterator struct {
	Event *IndexerAlterOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IndexerAlterOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexerAlterOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IndexerAlterOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IndexerAlterOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexerAlterOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexerAlterOwner represents a AlterOwner event raised by the Indexer contract.
type IndexerAlterOwner struct {
	Key  string
	Form common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlterOwner is a free log retrieval operation binding the contract event 0x46bd035a76a8302bb74520f9226b59925d8186784298f88ad636a4ea46b85b21.
//
// Solidity: event AlterOwner(string key, address form, address to)
func (_Indexer *IndexerFilterer) FilterAlterOwner(opts *bind.FilterOpts) (*IndexerAlterOwnerIterator, error) {

	logs, sub, err := _Indexer.contract.FilterLogs(opts, "AlterOwner")
	if err != nil {
		return nil, err
	}
	return &IndexerAlterOwnerIterator{contract: _Indexer.contract, event: "AlterOwner", logs: logs, sub: sub}, nil
}

// WatchAlterOwner is a free log subscription operation binding the contract event 0x46bd035a76a8302bb74520f9226b59925d8186784298f88ad636a4ea46b85b21.
//
// Solidity: event AlterOwner(string key, address form, address to)
func (_Indexer *IndexerFilterer) WatchAlterOwner(opts *bind.WatchOpts, sink chan<- *IndexerAlterOwner) (event.Subscription, error) {

	logs, sub, err := _Indexer.contract.WatchLogs(opts, "AlterOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexerAlterOwner)
				if err := _Indexer.contract.UnpackLog(event, "AlterOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IndexerAlterResolverIterator is returned from FilterAlterResolver and is used to iterate over the raw logs and unpacked data for AlterResolver events raised by the Indexer contract.
type IndexerAlterResolverIterator struct {
	Event *IndexerAlterResolver // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IndexerAlterResolverIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexerAlterResolver)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IndexerAlterResolver)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IndexerAlterResolverIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexerAlterResolverIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexerAlterResolver represents a AlterResolver event raised by the Indexer contract.
type IndexerAlterResolver struct {
	Key  string
	Form common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlterResolver is a free log retrieval operation binding the contract event 0x0a7047ba8be4d874e67aebc953a70ff6db03a81782549290ac646e0738ddfc04.
//
// Solidity: event AlterResolver(string key, address form, address to)
func (_Indexer *IndexerFilterer) FilterAlterResolver(opts *bind.FilterOpts) (*IndexerAlterResolverIterator, error) {

	logs, sub, err := _Indexer.contract.FilterLogs(opts, "AlterResolver")
	if err != nil {
		return nil, err
	}
	return &IndexerAlterResolverIterator{contract: _Indexer.contract, event: "AlterResolver", logs: logs, sub: sub}, nil
}

// WatchAlterResolver is a free log subscription operation binding the contract event 0x0a7047ba8be4d874e67aebc953a70ff6db03a81782549290ac646e0738ddfc04.
//
// Solidity: event AlterResolver(string key, address form, address to)
func (_Indexer *IndexerFilterer) WatchAlterResolver(opts *bind.WatchOpts, sink chan<- *IndexerAlterResolver) (event.Subscription, error) {

	logs, sub, err := _Indexer.contract.WatchLogs(opts, "AlterResolver")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexerAlterResolver)
				if err := _Indexer.contract.UnpackLog(event, "AlterResolver", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IndexerErrorIterator is returned from FilterError and is used to iterate over the raw logs and unpacked data for Error events raised by the Indexer contract.
type IndexerErrorIterator struct {
	Event *IndexerError // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IndexerErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IndexerError)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IndexerError)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IndexerErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IndexerErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IndexerError represents a Error event raised by the Indexer contract.
type IndexerError struct {
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterError is a free log retrieval operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string data)
func (_Indexer *IndexerFilterer) FilterError(opts *bind.FilterOpts) (*IndexerErrorIterator, error) {

	logs, sub, err := _Indexer.contract.FilterLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return &IndexerErrorIterator{contract: _Indexer.contract, event: "Error", logs: logs, sub: sub}, nil
}

// WatchError is a free log subscription operation binding the contract event 0x08c379a0afcc32b1a39302f7cb8073359698411ab5fd6e3edb2c02c0b5fba8aa.
//
// Solidity: event Error(string data)
func (_Indexer *IndexerFilterer) WatchError(opts *bind.WatchOpts, sink chan<- *IndexerError) (event.Subscription, error) {

	logs, sub, err := _Indexer.contract.WatchLogs(opts, "Error")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IndexerError)
				if err := _Indexer.contract.UnpackLog(event, "Error", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
