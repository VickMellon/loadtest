// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package MiniStore

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// MiniStoreABI is the input ABI used to generate the binding from.
const MiniStoreABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"addValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deleteLastArrayValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArrayDataLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getArrayLatestValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getArrayValueAt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_array\",\"type\":\"uint256[]\"}],\"name\":\"insertArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_i\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setArrayValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNumberValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MiniStoreBin is the compiled bytecode used for deploying new contracts.
var MiniStoreBin = "0x608060405234801561001057600080fd5b5061060f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c806369a127991161006657806369a12799146101fc57806375a01c0f1461021a57806399ff5f4814610248578063a366474f14610266578063fc7ef48c146102b257610093565b806314d0c25a146100985780631e25e5f8146100b657806356298b51146101785780635b9af12b146101ba575b600080fd5b6100a06102d0565b6040518082815260200191505060405180910390f35b610176600480360360408110156100cc57600080fd5b8101908080359060200190929190803590602001906401000000008111156100f357600080fd5b82018360208201111561010557600080fd5b8035906020019184602083028401116401000000008311171561012757600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050919291929050505061030e565b005b6101a46004803603602081101561018e57600080fd5b81019080803590602001909291905050506103b1565b6040518082815260200191505060405180910390f35b6101e6600480360360208110156101d057600080fd5b810190808035906020019092919050505061044b565b6040518082815260200191505060405180910390f35b610204610485565b6040518082815260200191505060405180910390f35b6102466004803603602081101561023057600080fd5b8101908080359060200190929190505050610491565b005b61025061049b565b6040518082815260200191505060405180910390f35b61029c6004803603604081101561027c57600080fd5b8101908080359060200190929190803590602001909291905050506104a5565b6040518082815260200191505060405180910390f35b6102ba610512565b6040518082815260200191505060405180910390f35b6000806000805490501115610306576000600160008054905003815481106102f457fe5b9060005260206000200154905061030b565b600090505b90565b60008090505b81518110156103ac5760008054905083820110156103625781818151811061033857fe5b602002602001015160008483018154811061034f57fe5b906000526020600020018190555061039f565b600082828151811061037057fe5b602002602001015190806001815401808255809150506001900390600052602060002001600090919091909150555b8080600101915050610314565b505050565b60008080549050821061042c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f496e646578206f7574206f662072616e6765000000000000000000000000000081525060200191505060405180910390fd5b6000828154811061043957fe5b90600052602060002001549050919050565b6000808290806001815401808255809150506001900390600052602060002001600090919091909150556001600080549050039050919050565b60008080549050905090565b8060018190555050565b6000600154905090565b600080805490508310156104d75781600084815481106104c157fe5b906000526020600020018190555082905061050c565b600082908060018154018082558091505060019003906000526020600020016000909190919091505560016000805490500390505b92915050565b6000806000805490501161058e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600e8152602001807f617272617920697320656d70747900000000000000000000000000000000000081525060200191505060405180910390fd5b600080600160008054905003815481106105a457fe5b9060005260206000200154905060008054806105bc57fe5b60019003818190600052602060002001600090559055809150509056fea2646970667358221220df75a650977e533c700991fc145b1db8db69559d46fabe4b662f814dea7a8bd964736f6c63430006000033"

// DeployMiniStore deploys a new Ethereum contract, binding an instance of MiniStore to it.
func DeployMiniStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MiniStore, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniStoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MiniStoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiniStore{MiniStoreCaller: MiniStoreCaller{contract: contract}, MiniStoreTransactor: MiniStoreTransactor{contract: contract}, MiniStoreFilterer: MiniStoreFilterer{contract: contract}}, nil
}

// MiniStore is an auto generated Go binding around an Ethereum contract.
type MiniStore struct {
	MiniStoreCaller     // Read-only binding to the contract
	MiniStoreTransactor // Write-only binding to the contract
	MiniStoreFilterer   // Log filterer for contract events
}

// MiniStoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiniStoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniStoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiniStoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniStoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiniStoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiniStoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiniStoreSession struct {
	Contract     *MiniStore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiniStoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiniStoreCallerSession struct {
	Contract *MiniStoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MiniStoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiniStoreTransactorSession struct {
	Contract     *MiniStoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MiniStoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiniStoreRaw struct {
	Contract *MiniStore // Generic contract binding to access the raw methods on
}

// MiniStoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiniStoreCallerRaw struct {
	Contract *MiniStoreCaller // Generic read-only contract binding to access the raw methods on
}

// MiniStoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiniStoreTransactorRaw struct {
	Contract *MiniStoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiniStore creates a new instance of MiniStore, bound to a specific deployed contract.
func NewMiniStore(address common.Address, backend bind.ContractBackend) (*MiniStore, error) {
	contract, err := bindMiniStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiniStore{MiniStoreCaller: MiniStoreCaller{contract: contract}, MiniStoreTransactor: MiniStoreTransactor{contract: contract}, MiniStoreFilterer: MiniStoreFilterer{contract: contract}}, nil
}

// NewMiniStoreCaller creates a new read-only instance of MiniStore, bound to a specific deployed contract.
func NewMiniStoreCaller(address common.Address, caller bind.ContractCaller) (*MiniStoreCaller, error) {
	contract, err := bindMiniStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiniStoreCaller{contract: contract}, nil
}

// NewMiniStoreTransactor creates a new write-only instance of MiniStore, bound to a specific deployed contract.
func NewMiniStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*MiniStoreTransactor, error) {
	contract, err := bindMiniStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiniStoreTransactor{contract: contract}, nil
}

// NewMiniStoreFilterer creates a new log filterer instance of MiniStore, bound to a specific deployed contract.
func NewMiniStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*MiniStoreFilterer, error) {
	contract, err := bindMiniStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiniStoreFilterer{contract: contract}, nil
}

// bindMiniStore binds a generic wrapper to an already deployed contract.
func bindMiniStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MiniStoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniStore *MiniStoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiniStore.Contract.MiniStoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniStore *MiniStoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniStore.Contract.MiniStoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniStore *MiniStoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniStore.Contract.MiniStoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiniStore *MiniStoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiniStore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiniStore *MiniStoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniStore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiniStore *MiniStoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiniStore.Contract.contract.Transact(opts, method, params...)
}

// GetArrayDataLength is a free data retrieval call binding the contract method 0x69a12799.
//
// Solidity: function getArrayDataLength() view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetArrayDataLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getArrayDataLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArrayDataLength is a free data retrieval call binding the contract method 0x69a12799.
//
// Solidity: function getArrayDataLength() view returns(uint256)
func (_MiniStore *MiniStoreSession) GetArrayDataLength() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayDataLength(&_MiniStore.CallOpts)
}

// GetArrayDataLength is a free data retrieval call binding the contract method 0x69a12799.
//
// Solidity: function getArrayDataLength() view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetArrayDataLength() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayDataLength(&_MiniStore.CallOpts)
}

// GetArrayLatestValue is a free data retrieval call binding the contract method 0x14d0c25a.
//
// Solidity: function getArrayLatestValue() view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetArrayLatestValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getArrayLatestValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArrayLatestValue is a free data retrieval call binding the contract method 0x14d0c25a.
//
// Solidity: function getArrayLatestValue() view returns(uint256)
func (_MiniStore *MiniStoreSession) GetArrayLatestValue() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayLatestValue(&_MiniStore.CallOpts)
}

// GetArrayLatestValue is a free data retrieval call binding the contract method 0x14d0c25a.
//
// Solidity: function getArrayLatestValue() view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetArrayLatestValue() (*big.Int, error) {
	return _MiniStore.Contract.GetArrayLatestValue(&_MiniStore.CallOpts)
}

// GetArrayValueAt is a free data retrieval call binding the contract method 0x56298b51.
//
// Solidity: function getArrayValueAt(uint256 i) view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetArrayValueAt(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getArrayValueAt", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetArrayValueAt is a free data retrieval call binding the contract method 0x56298b51.
//
// Solidity: function getArrayValueAt(uint256 i) view returns(uint256)
func (_MiniStore *MiniStoreSession) GetArrayValueAt(i *big.Int) (*big.Int, error) {
	return _MiniStore.Contract.GetArrayValueAt(&_MiniStore.CallOpts, i)
}

// GetArrayValueAt is a free data retrieval call binding the contract method 0x56298b51.
//
// Solidity: function getArrayValueAt(uint256 i) view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetArrayValueAt(i *big.Int) (*big.Int, error) {
	return _MiniStore.Contract.GetArrayValueAt(&_MiniStore.CallOpts, i)
}

// GetNumberValue is a free data retrieval call binding the contract method 0x99ff5f48.
//
// Solidity: function getNumberValue() view returns(uint256)
func (_MiniStore *MiniStoreCaller) GetNumberValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MiniStore.contract.Call(opts, &out, "getNumberValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberValue is a free data retrieval call binding the contract method 0x99ff5f48.
//
// Solidity: function getNumberValue() view returns(uint256)
func (_MiniStore *MiniStoreSession) GetNumberValue() (*big.Int, error) {
	return _MiniStore.Contract.GetNumberValue(&_MiniStore.CallOpts)
}

// GetNumberValue is a free data retrieval call binding the contract method 0x99ff5f48.
//
// Solidity: function getNumberValue() view returns(uint256)
func (_MiniStore *MiniStoreCallerSession) GetNumberValue() (*big.Int, error) {
	return _MiniStore.Contract.GetNumberValue(&_MiniStore.CallOpts)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactor) AddValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "addValue", _value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreSession) AddValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.AddValue(&_MiniStore.TransactOpts, _value)
}

// AddValue is a paid mutator transaction binding the contract method 0x5b9af12b.
//
// Solidity: function addValue(uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactorSession) AddValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.AddValue(&_MiniStore.TransactOpts, _value)
}

// DeleteLastArrayValue is a paid mutator transaction binding the contract method 0xfc7ef48c.
//
// Solidity: function deleteLastArrayValue() returns(uint256)
func (_MiniStore *MiniStoreTransactor) DeleteLastArrayValue(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "deleteLastArrayValue")
}

// DeleteLastArrayValue is a paid mutator transaction binding the contract method 0xfc7ef48c.
//
// Solidity: function deleteLastArrayValue() returns(uint256)
func (_MiniStore *MiniStoreSession) DeleteLastArrayValue() (*types.Transaction, error) {
	return _MiniStore.Contract.DeleteLastArrayValue(&_MiniStore.TransactOpts)
}

// DeleteLastArrayValue is a paid mutator transaction binding the contract method 0xfc7ef48c.
//
// Solidity: function deleteLastArrayValue() returns(uint256)
func (_MiniStore *MiniStoreTransactorSession) DeleteLastArrayValue() (*types.Transaction, error) {
	return _MiniStore.Contract.DeleteLastArrayValue(&_MiniStore.TransactOpts)
}

// InsertArray is a paid mutator transaction binding the contract method 0x1e25e5f8.
//
// Solidity: function insertArray(uint256 _startIndex, uint256[] _array) returns()
func (_MiniStore *MiniStoreTransactor) InsertArray(opts *bind.TransactOpts, _startIndex *big.Int, _array []*big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "insertArray", _startIndex, _array)
}

// InsertArray is a paid mutator transaction binding the contract method 0x1e25e5f8.
//
// Solidity: function insertArray(uint256 _startIndex, uint256[] _array) returns()
func (_MiniStore *MiniStoreSession) InsertArray(_startIndex *big.Int, _array []*big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.InsertArray(&_MiniStore.TransactOpts, _startIndex, _array)
}

// InsertArray is a paid mutator transaction binding the contract method 0x1e25e5f8.
//
// Solidity: function insertArray(uint256 _startIndex, uint256[] _array) returns()
func (_MiniStore *MiniStoreTransactorSession) InsertArray(_startIndex *big.Int, _array []*big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.InsertArray(&_MiniStore.TransactOpts, _startIndex, _array)
}

// SetArrayValue is a paid mutator transaction binding the contract method 0xa366474f.
//
// Solidity: function setArrayValue(uint256 _i, uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactor) SetArrayValue(opts *bind.TransactOpts, _i *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "setArrayValue", _i, _value)
}

// SetArrayValue is a paid mutator transaction binding the contract method 0xa366474f.
//
// Solidity: function setArrayValue(uint256 _i, uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreSession) SetArrayValue(_i *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetArrayValue(&_MiniStore.TransactOpts, _i, _value)
}

// SetArrayValue is a paid mutator transaction binding the contract method 0xa366474f.
//
// Solidity: function setArrayValue(uint256 _i, uint256 _value) returns(uint256)
func (_MiniStore *MiniStoreTransactorSession) SetArrayValue(_i *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetArrayValue(&_MiniStore.TransactOpts, _i, _value)
}

// SetNumberValue is a paid mutator transaction binding the contract method 0x75a01c0f.
//
// Solidity: function setNumberValue(uint256 _value) returns()
func (_MiniStore *MiniStoreTransactor) SetNumberValue(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _MiniStore.contract.Transact(opts, "setNumberValue", _value)
}

// SetNumberValue is a paid mutator transaction binding the contract method 0x75a01c0f.
//
// Solidity: function setNumberValue(uint256 _value) returns()
func (_MiniStore *MiniStoreSession) SetNumberValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetNumberValue(&_MiniStore.TransactOpts, _value)
}

// SetNumberValue is a paid mutator transaction binding the contract method 0x75a01c0f.
//
// Solidity: function setNumberValue(uint256 _value) returns()
func (_MiniStore *MiniStoreTransactorSession) SetNumberValue(_value *big.Int) (*types.Transaction, error) {
	return _MiniStore.Contract.SetNumberValue(&_MiniStore.TransactOpts, _value)
}
