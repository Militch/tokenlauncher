// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// ERC20TokenContractABI is the input ABI used to generate the binding from.
const ERC20TokenContractABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_initailSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20TokenContractBin is the compiled bytecode used for deploying new contracts.
var ERC20TokenContractBin = "0x60806040523480156200001157600080fd5b5060405162000d5e38038062000d5e83398101604081905262000034916200026a565b83516200004990600090602087019062000111565b5082516200005f90600190602086019062000111565b506002805460ff191660ff8381169190911791829055620000839116600a62000340565b6200008f908362000438565b6003819055600480546001600160a01b03191633178082556001600160a01b0390811660009081526005602052604080822085905592549251929091169290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91620000ff9190815260200190565b60405180910390a350505050620004c3565b8280546200011f906200045a565b90600052602060002090601f0160209004810192826200014357600085556200018e565b82601f106200015e57805160ff19168380011785556200018e565b828001600101855582156200018e579182015b828111156200018e57825182559160200191906001019062000171565b506200019c929150620001a0565b5090565b5b808211156200019c5760008155600101620001a1565b600082601f830112620001c8578081fd5b81516001600160401b0380821115620001e557620001e5620004ad565b604051601f8301601f19908116603f01168101908282118183101715620002105762000210620004ad565b816040528381526020925086838588010111156200022c578485fd5b8491505b838210156200024f578582018301518183018401529082019062000230565b838211156200026057848385830101525b9695505050505050565b6000806000806080858703121562000280578384fd5b84516001600160401b038082111562000297578586fd5b620002a588838901620001b7565b95506020870151915080821115620002bb578485fd5b50620002ca87828801620001b7565b93505060408501519150606085015160ff81168114620002e8578182fd5b939692955090935050565b80825b600180861162000307575062000337565b8187048211156200031c576200031c62000497565b808616156200032a57918102915b9490941c938002620002f6565b94509492505050565b60006200035460001960ff8516846200035b565b9392505050565b6000826200036c5750600162000354565b816200037b5750600062000354565b81600181146200039457600281146200039f57620003d3565b600191505062000354565b60ff841115620003b357620003b362000497565b6001841b915084821115620003cc57620003cc62000497565b5062000354565b5060208310610133831016604e8410600b84101617156200040b575081810a8381111562000405576200040562000497565b62000354565b6200041a8484846001620002f3565b8086048211156200042f576200042f62000497565b02949350505050565b600081600019048311821515161562000455576200045562000497565b500290565b6002810460018216806200046f57607f821691505b602082108114156200049157634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b61088b80620004d36000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063313ce56711610066578063313ce5671461013457806370a082311461015357806395d89b4114610166578063a9059cbb1461016e578063dd62ed3e146101815761009e565b806306fdde03146100a357806307546172146100c1578063095ea7b3146100ec57806318160ddd1461010f57806323b872dd14610121575b600080fd5b6100ab6101ba565b6040516100b89190610782565b60405180910390f35b6004546100d4906001600160a01b031681565b6040516001600160a01b0390911681526020016100b8565b6100ff6100fa366004610759565b610248565b60405190151581526020016100b8565b6003545b6040519081526020016100b8565b6100ff61012f36600461071e565b6102b4565b6002546101419060ff1681565b60405160ff90911681526020016100b8565b6101136101613660046106d2565b6104f1565b6100ab610510565b6100ff61017c366004610759565b61051d565b61011361018f3660046106ec565b6001600160a01b03918216600090815260066020908152604080832093909416825291909152205490565b600080546101c790610804565b80601f01602080910402602001604051908101604052809291908181526020018280546101f390610804565b80156102405780601f1061021557610100808354040283529160200191610240565b820191906000526020600020905b81548152906001019060200180831161022357829003601f168201915b505050505081565b3360008181526006602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906102a39086815260200190565b60405180910390a350600192915050565b6001600160a01b0383166000908152600560205260408120548211156103215760405162461bcd60e51b815260206004820181905260248201527f66726f6d20646f65736e74206861766520656e6f7567682062616c616e63652e60448201526064015b60405180910390fd5b6001600160a01b03841660009081526006602090815260408083203384529091529020548211156103a35760405162461bcd60e51b815260206004820152602660248201527f416c6c6f77616e6365206f66206d73672e73656e646572206973206e6f7420656044820152653737bab3b41760d11b6064820152608401610318565b6001600160a01b0383166103f95760405162461bcd60e51b815260206004820181905260248201527f43616e6e6f742073656e6420746f20616c6c207a65726f20616464726573732e6044820152606401610318565b6001600160a01b03841660009081526005602052604090205461041c9083610675565b6001600160a01b03808616600090815260056020526040808220939093559085168152205461044b9083610698565b6001600160a01b0380851660009081526005602090815260408083209490945591871681526006825282812033825290915220546104899083610675565b6001600160a01b03858116600081815260066020908152604080832033845282529182902094909455518581529186169290917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35060019392505050565b6001600160a01b0381166000908152600560205260409020545b919050565b600180546101c790610804565b60006001600160a01b0383166105755760405162461bcd60e51b815260206004820181905260248201527f43616e6e6f742073656e6420746f20616c6c207a65726f20616464726573732e6044820152606401610318565b336000908152600560205260409020548211156105de5760405162461bcd60e51b815260206004820152602160248201527f6d73672e73656e6465722062616c616e6365206973206e6f7420656e6f7567686044820152601760f91b6064820152608401610318565b336000908152600560205260409020546105f89083610675565b33600090815260056020526040808220929092556001600160a01b038516815220546106249083610698565b6001600160a01b0384166000818152600560205260409081902092909255905133907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906102a39086815260200190565b60008282111561068457600080fd5b600061069083856107ed565b949350505050565b6000806106a583856107d5565b9050838110156106b457600080fd5b9392505050565b80356001600160a01b038116811461050b57600080fd5b6000602082840312156106e3578081fd5b6106b4826106bb565b600080604083850312156106fe578081fd5b610707836106bb565b9150610715602084016106bb565b90509250929050565b600080600060608486031215610732578081fd5b61073b846106bb565b9250610749602085016106bb565b9150604084013590509250925092565b6000806040838503121561076b578182fd5b610774836106bb565b946020939093013593505050565b6000602080835283518082850152825b818110156107ae57858101830151858201604001528201610792565b818111156107bf5783604083870101525b50601f01601f1916929092016040019392505050565b600082198211156107e8576107e861083f565b500190565b6000828210156107ff576107ff61083f565b500390565b60028104600182168061081857607f821691505b6020821081141561083957634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220b5024a803e45a1b8860ac26aa665ada88735f9e582a64d0ff59e5d06dd2e09fe64736f6c63430008020033"

// DeployERC20TokenContract deploys a new Ethereum contract, binding an instance of ERC20TokenContract to it.
func DeployERC20TokenContract(auth *bind.TransactOpts, backend bind.ContractBackend, _name string, _symbol string, _initailSupply *big.Int, _decimals uint8) (common.Address, *types.Transaction, *ERC20TokenContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20TokenContractBin), backend, _name, _symbol, _initailSupply, _decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20TokenContract{ERC20TokenContractCaller: ERC20TokenContractCaller{contract: contract}, ERC20TokenContractTransactor: ERC20TokenContractTransactor{contract: contract}, ERC20TokenContractFilterer: ERC20TokenContractFilterer{contract: contract}}, nil
}

// ERC20TokenContract is an auto generated Go binding around an Ethereum contract.
type ERC20TokenContract struct {
	ERC20TokenContractCaller     // Read-only binding to the contract
	ERC20TokenContractTransactor // Write-only binding to the contract
	ERC20TokenContractFilterer   // Log filterer for contract events
}

// ERC20TokenContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20TokenContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20TokenContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20TokenContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20TokenContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20TokenContractSession struct {
	Contract     *ERC20TokenContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ERC20TokenContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20TokenContractCallerSession struct {
	Contract *ERC20TokenContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ERC20TokenContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TokenContractTransactorSession struct {
	Contract     *ERC20TokenContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ERC20TokenContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20TokenContractRaw struct {
	Contract *ERC20TokenContract // Generic contract binding to access the raw methods on
}

// ERC20TokenContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20TokenContractCallerRaw struct {
	Contract *ERC20TokenContractCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20TokenContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TokenContractTransactorRaw struct {
	Contract *ERC20TokenContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20TokenContract creates a new instance of ERC20TokenContract, bound to a specific deployed contract.
func NewERC20TokenContract(address common.Address, backend bind.ContractBackend) (*ERC20TokenContract, error) {
	contract, err := bindERC20TokenContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenContract{ERC20TokenContractCaller: ERC20TokenContractCaller{contract: contract}, ERC20TokenContractTransactor: ERC20TokenContractTransactor{contract: contract}, ERC20TokenContractFilterer: ERC20TokenContractFilterer{contract: contract}}, nil
}

// NewERC20TokenContractCaller creates a new read-only instance of ERC20TokenContract, bound to a specific deployed contract.
func NewERC20TokenContractCaller(address common.Address, caller bind.ContractCaller) (*ERC20TokenContractCaller, error) {
	contract, err := bindERC20TokenContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenContractCaller{contract: contract}, nil
}

// NewERC20TokenContractTransactor creates a new write-only instance of ERC20TokenContract, bound to a specific deployed contract.
func NewERC20TokenContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20TokenContractTransactor, error) {
	contract, err := bindERC20TokenContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenContractTransactor{contract: contract}, nil
}

// NewERC20TokenContractFilterer creates a new log filterer instance of ERC20TokenContract, bound to a specific deployed contract.
func NewERC20TokenContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20TokenContractFilterer, error) {
	contract, err := bindERC20TokenContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenContractFilterer{contract: contract}, nil
}

// bindERC20TokenContract binds a generic wrapper to an already deployed contract.
func bindERC20TokenContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20TokenContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenContract *ERC20TokenContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenContract.Contract.ERC20TokenContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenContract *ERC20TokenContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.ERC20TokenContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenContract *ERC20TokenContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.ERC20TokenContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20TokenContract *ERC20TokenContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20TokenContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20TokenContract *ERC20TokenContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20TokenContract *ERC20TokenContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenContract.Contract.Allowance(&_ERC20TokenContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20TokenContract.Contract.Allowance(&_ERC20TokenContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20TokenContract.Contract.BalanceOf(&_ERC20TokenContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC20TokenContract.Contract.BalanceOf(&_ERC20TokenContract.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenContract *ERC20TokenContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenContract *ERC20TokenContractSession) Decimals() (uint8, error) {
	return _ERC20TokenContract.Contract.Decimals(&_ERC20TokenContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) Decimals() (uint8, error) {
	return _ERC20TokenContract.Contract.Decimals(&_ERC20TokenContract.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20TokenContract *ERC20TokenContractCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "minter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20TokenContract *ERC20TokenContractSession) Minter() (common.Address, error) {
	return _ERC20TokenContract.Contract.Minter(&_ERC20TokenContract.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() view returns(address)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) Minter() (common.Address, error) {
	return _ERC20TokenContract.Contract.Minter(&_ERC20TokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenContract *ERC20TokenContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenContract *ERC20TokenContractSession) Name() (string, error) {
	return _ERC20TokenContract.Contract.Name(&_ERC20TokenContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) Name() (string, error) {
	return _ERC20TokenContract.Contract.Name(&_ERC20TokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenContract *ERC20TokenContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenContract *ERC20TokenContractSession) Symbol() (string, error) {
	return _ERC20TokenContract.Contract.Symbol(&_ERC20TokenContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) Symbol() (string, error) {
	return _ERC20TokenContract.Contract.Symbol(&_ERC20TokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20TokenContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenContract.Contract.TotalSupply(&_ERC20TokenContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20TokenContract *ERC20TokenContractCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20TokenContract.Contract.TotalSupply(&_ERC20TokenContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.Approve(&_ERC20TokenContract.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.Approve(&_ERC20TokenContract.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.Transfer(&_ERC20TokenContract.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.Transfer(&_ERC20TokenContract.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.TransferFrom(&_ERC20TokenContract.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_ERC20TokenContract *ERC20TokenContractTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC20TokenContract.Contract.TransferFrom(&_ERC20TokenContract.TransactOpts, from, to, value)
}

// ERC20TokenContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20TokenContract contract.
type ERC20TokenContractApprovalIterator struct {
	Event *ERC20TokenContractApproval // Event containing the contract specifics and raw log

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
func (it *ERC20TokenContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenContractApproval)
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
		it.Event = new(ERC20TokenContractApproval)
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
func (it *ERC20TokenContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenContractApproval represents a Approval event raised by the ERC20TokenContract contract.
type ERC20TokenContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenContract *ERC20TokenContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20TokenContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenContractApprovalIterator{contract: _ERC20TokenContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenContract *ERC20TokenContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20TokenContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20TokenContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenContractApproval)
				if err := _ERC20TokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20TokenContract *ERC20TokenContractFilterer) ParseApproval(log types.Log) (*ERC20TokenContractApproval, error) {
	event := new(ERC20TokenContractApproval)
	if err := _ERC20TokenContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20TokenContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20TokenContract contract.
type ERC20TokenContractTransferIterator struct {
	Event *ERC20TokenContractTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20TokenContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20TokenContractTransfer)
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
		it.Event = new(ERC20TokenContractTransfer)
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
func (it *ERC20TokenContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TokenContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20TokenContractTransfer represents a Transfer event raised by the ERC20TokenContract contract.
type ERC20TokenContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenContract *ERC20TokenContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TokenContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TokenContractTransferIterator{contract: _ERC20TokenContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenContract *ERC20TokenContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20TokenContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20TokenContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20TokenContractTransfer)
				if err := _ERC20TokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20TokenContract *ERC20TokenContractFilterer) ParseTransfer(log types.Log) (*ERC20TokenContractTransfer, error) {
	event := new(ERC20TokenContractTransfer)
	if err := _ERC20TokenContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
