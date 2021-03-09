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

// CrowdSaleABI is the input ABI used to generate the binding from.
const CrowdSaleABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetFunds\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_endTime\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raise\",\"type\":\"uint256\"}],\"name\":\"Bought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raise\",\"type\":\"uint256\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"raise\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsRaisedTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fundsSoldTotal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFinished\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSalesCompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUncompleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"price\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setEndTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// CrowdSaleFuncSigs maps the 4-byte function signature to its string representation.
var CrowdSaleFuncSigs = map[string]string{
	"a6f2ae3a": "buy()",
	"3197cbb6": "endTime()",
	"eeeb3650": "fundsRaisedTotal()",
	"db8c69ef": "fundsSoldTotal()",
	"22f3e2d4": "isActive()",
	"7b352962": "isFinished()",
	"a06c944a": "isSalesCompleted()",
	"e2481cce": "isUncompleted()",
	"8da5cb5b": "owner()",
	"a035b1fe": "price()",
	"590e1ae3": "refund()",
	"ccb98ffc": "setEndTime(uint256)",
	"91b7f5ed": "setPrice(uint256)",
	"3e0a322d": "setStartTime(uint256)",
	"78e97925": "startTime()",
	"1459e2ad": "targetFunds()",
	"fc0c546a": "token()",
	"1a39d8ef": "totalAmount()",
	"f2fde38b": "transferOwnership(address)",
	"3ccfd60b": "withdraw()",
}

// CrowdSaleBin is the compiled bytecode used for deploying new contracts.
var CrowdSaleBin = "0x608060405260006001556000600255600060045560006005553480156200002557600080fd5b506040516200129d3803806200129d83398101604081905262000048916200024c565b600080546001600160a01b03199081163317909155600380546001600160a01b038881169190931617908190556040805163313ce56760e01b81529051919092169163313ce567916004808301926020929190829003018186803b158015620000b057600080fd5b505afa158015620000c5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190620000eb9190620002a1565b620000fb9060ff16600a62000318565b62000107908562000400565b600555600483905581620001715760405162461bcd60e51b815260206004820152602660248201527f737461727454696d65206d7573742062652067726561746572207468616e20746044820152656f207a65726f60d01b60648201526084015b60405180910390fd5b60008111620001cf5760405162461bcd60e51b8152602060048201526024808201527f656e6454696d65206d7573742062652067726561746572207468616e20746f206044820152637a65726f60e01b606482015260840162000168565b80821115620002395760405162461bcd60e51b815260206004820152602f60248201527f737461727454696d65206d757374206265206c657373207468616e206f72206560448201526e7175616c20746f20656e6454696d6560881b606482015260840162000168565b6006919091556007555062000438915050565b600080600080600060a0868803121562000264578081fd5b85516001600160a01b03811681146200027b578182fd5b602087015160408801516060890151608090990151929a91995097965090945092505050565b600060208284031215620002b3578081fd5b815160ff81168114620002c4578182fd5b9392505050565b80825b6001808611620002df57506200030f565b818704821115620002f457620002f462000422565b808616156200030257918102915b9490941c938002620002ce565b94509492505050565b6000620002c460001984846000826200033457506001620002c4565b816200034357506000620002c4565b81600181146200035c576002811462000367576200039b565b6001915050620002c4565b60ff8411156200037b576200037b62000422565b6001841b91508482111562000394576200039462000422565b50620002c4565b5060208310610133831016604e8410600b8410161715620003d3575081810a83811115620003cd57620003cd62000422565b620002c4565b620003e28484846001620002cb565b808604821115620003f757620003f762000422565b02949350505050565b60008160001904831182151516156200041d576200041d62000422565b500290565b634e487b7160e01b600052601160045260246000fd5b610e5580620004486000396000f3fe6080604052600436106101235760003560e01c806391b7f5ed116100a0578063db8c69ef11610064578063db8c69ef146102df578063e2481cce146102f5578063eeeb36501461030a578063f2fde38b14610320578063fc0c546a1461034057610132565b806391b7f5ed14610267578063a035b1fe14610287578063a06c944a1461029d578063a6f2ae3a146102b7578063ccb98ffc146102bf57610132565b80633e0a322d116100e75780633e0a322d146101cf578063590e1ae3146101ef57806378e97925146102045780637b3529621461021a5780638da5cb5b1461022f57610132565b80631459e2ad146101415780631a39d8ef1461016a57806322f3e2d41461017f5780633197cbb6146101a45780633ccfd60b146101ba57610132565b3661013257610130610360565b005b34801561013e57600080fd5b50005b34801561014d57600080fd5b5061015760055481565b6040519081526020015b60405180910390f35b34801561017657600080fd5b506101576106be565b34801561018b57600080fd5b5061019461073f565b6040519015158152602001610161565b3480156101b057600080fd5b5061015760075481565b3480156101c657600080fd5b50610194610770565b3480156101db57600080fd5b506101946101ea366004610c16565b610876565b3480156101fb57600080fd5b506101946108aa565b34801561021057600080fd5b5061015760065481565b34801561022657600080fd5b50610194610a04565b34801561023b57600080fd5b5060005461024f906001600160a01b031681565b6040516001600160a01b039091168152602001610161565b34801561027357600080fd5b50610194610282366004610c16565b610a22565b34801561029357600080fd5b5061015760045481565b3480156102a957600080fd5b506005546001541015610194565b610130610360565b3480156102cb57600080fd5b506101946102da366004610c16565b610a56565b3480156102eb57600080fd5b5061015760015481565b34801561030157600080fd5b50610194610a8a565b34801561031657600080fd5b5061015760025481565b34801561032c57600080fd5b5061013061033b366004610bcf565b610aa8565b34801561034c57600080fd5b5060035461024f906001600160a01b031681565b61036861073f565b6103b95760405162461bcd60e51b815260206004820152601860248201527f43726f77642073616c65206973206e6f7420616374697665000000000000000060448201526064015b60405180910390fd5b6003546040805163313ce56760e01b8152905134926000926001600160a01b039091169163313ce56791600480820192602092909190829003018186803b15801561040357600080fd5b505afa158015610417573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043b9190610c46565b60ff16905061046261044e82600a610d1c565b60045461045c908590610b56565b90610b7a565b6003546040516370a0823160e01b81523060048201529193506000916001600160a01b03909116906370a082319060240160206040518083038186803b1580156104ab57600080fd5b505afa1580156104bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104e39190610c2e565b9050600083116105355760405162461bcd60e51b815260206004820152601b60248201527f596f75206e65656420746f2073656e6420736f6d65206574686572000000000060448201526064016103b0565b808311156105855760405162461bcd60e51b815260206004820181905260248201527f4e6f7420656e6f75676820746f6b656e7320696e20746865207265736572766560448201526064016103b0565b60035460405163a9059cbb60e01b8152336004820152602481018590526001600160a01b039091169063a9059cbb90604401602060405180830381600087803b1580156105d157600080fd5b505af11580156105e5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106099190610bf6565b506001546106179084610bb3565b6001556002546106279034610bb3565b600255336000908152600860205260409020546106449034610bb3565b3360009081526008602090815260408083209390935560099052205461066a9084610bb3565b3360008181526009602090815260409182902093909355805134815292830186905290917fa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c910160405180910390a2505050565b6003546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a082319060240160206040518083038186803b15801561070257600080fd5b505afa158015610716573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061073a9190610c2e565b905090565b6006546000904290811080159061075857506007548111155b801561076a575060055460015410155b155b91505090565b600080546001600160a01b0316331461079b5760405162461bcd60e51b81526004016103b090610c67565b6107a3610a04565b6107ef5760405162461bcd60e51b815260206004820152601a60248201527f43726f77642073616c65206973206e6f742066696e697368656400000000000060448201526064016103b0565b600080546002546040516001600160a01b039092169281156108fc029290818181858888f1935050505015801561082a573d6000803e3d6000fd5b506000546002546040519081526001600160a01b03909116907f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d59060200160405180910390a250600190565b600080546001600160a01b031633146108a15760405162461bcd60e51b81526004016103b090610c67565b50600655600190565b60006108b4610a8a565b6109005760405162461bcd60e51b815260206004820152601960248201527f43726f77642073616c6520697320756e636f6d706c657465640000000000000060448201526064016103b0565b33600090815260086020908152604080832054600990925290912054816109775760405162461bcd60e51b815260206004820152602560248201527f596f75206e65656420746f2073656c6c206174206c6561737420736f6d6520746044820152646f6b656e7360d81b60648201526084016103b0565b604051339083156108fc029084906000818181858888f193505050501580156109a4573d6000803e3d6000fd5b50336000818152600860209081526040808320839055600982528083209290925581518581529081018490527f2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66910160405180910390a260019250505090565b60075460009042908111801561076a5750600554600154101561076a565b600080546001600160a01b03163314610a4d5760405162461bcd60e51b81526004016103b090610c67565b50600455600190565b600080546001600160a01b03163314610a815760405162461bcd60e51b81526004016103b090610c67565b50600755600190565b60075460009042908111801561076a57506005546001541015610768565b6000546001600160a01b03163314610ad25760405162461bcd60e51b81526004016103b090610c67565b6001600160a01b038116610b345760405162461bcd60e51b8152602060048201526024808201527f43616e6e6f74207472616e7366657220746f20616c6c207a65726f206164647260448201526332b9b99760e11b60648201526084016103b0565b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000808211610b6457600080fd5b6000610b708385610cb6565b9150505b92915050565b600082610b8957506000610b74565b6000610b958385610dea565b905082610ba28583610cb6565b14610bac57600080fd5b9392505050565b600080610bc08385610c9e565b905083811015610bac57600080fd5b600060208284031215610be0578081fd5b81356001600160a01b0381168114610bac578182fd5b600060208284031215610c07578081fd5b81518015158114610bac578182fd5b600060208284031215610c27578081fd5b5035919050565b600060208284031215610c3f578081fd5b5051919050565b600060208284031215610c57578081fd5b815160ff81168114610bac578182fd5b60208082526019908201527f4f6e6c79206f776e65722063616e2063616c6c20746869732e00000000000000604082015260600190565b60008219821115610cb157610cb1610e09565b500190565b600082610cd157634e487b7160e01b81526012600452602481fd5b500490565b80825b6001808611610ce85750610d13565b818704821115610cfa57610cfa610e09565b80861615610d0757918102915b9490941c938002610cd9565b94509492505050565b6000610bac6000198484600082610d3557506001610bac565b81610d4257506000610bac565b8160018114610d585760028114610d6257610d8f565b6001915050610bac565b60ff841115610d7357610d73610e09565b6001841b915084821115610d8957610d89610e09565b50610bac565b5060208310610133831016604e8410600b8410161715610dc2575081810a83811115610dbd57610dbd610e09565b610bac565b610dcf8484846001610cd6565b808604821115610de157610de1610e09565b02949350505050565b6000816000190483118215151615610e0457610e04610e09565b500290565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220524795772e023f0a235a87c10f6453b7fda9011ba52aa91ec7a11b2c676998b664736f6c63430008020033"

// DeployCrowdSale deploys a new Ethereum contract, binding an instance of CrowdSale to it.
func DeployCrowdSale(auth *bind.TransactOpts, backend bind.ContractBackend, tokenAddress common.Address, _targetFunds *big.Int, _price *big.Int, _startTime *big.Int, _endTime *big.Int) (common.Address, *types.Transaction, *CrowdSale, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdSaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CrowdSaleBin), backend, tokenAddress, _targetFunds, _price, _startTime, _endTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CrowdSale{CrowdSaleCaller: CrowdSaleCaller{contract: contract}, CrowdSaleTransactor: CrowdSaleTransactor{contract: contract}, CrowdSaleFilterer: CrowdSaleFilterer{contract: contract}}, nil
}

// CrowdSale is an auto generated Go binding around an Ethereum contract.
type CrowdSale struct {
	CrowdSaleCaller     // Read-only binding to the contract
	CrowdSaleTransactor // Write-only binding to the contract
	CrowdSaleFilterer   // Log filterer for contract events
}

// CrowdSaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrowdSaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdSaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrowdSaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdSaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrowdSaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrowdSaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrowdSaleSession struct {
	Contract     *CrowdSale        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrowdSaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrowdSaleCallerSession struct {
	Contract *CrowdSaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrowdSaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrowdSaleTransactorSession struct {
	Contract     *CrowdSaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrowdSaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrowdSaleRaw struct {
	Contract *CrowdSale // Generic contract binding to access the raw methods on
}

// CrowdSaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrowdSaleCallerRaw struct {
	Contract *CrowdSaleCaller // Generic read-only contract binding to access the raw methods on
}

// CrowdSaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrowdSaleTransactorRaw struct {
	Contract *CrowdSaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrowdSale creates a new instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSale(address common.Address, backend bind.ContractBackend) (*CrowdSale, error) {
	contract, err := bindCrowdSale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrowdSale{CrowdSaleCaller: CrowdSaleCaller{contract: contract}, CrowdSaleTransactor: CrowdSaleTransactor{contract: contract}, CrowdSaleFilterer: CrowdSaleFilterer{contract: contract}}, nil
}

// NewCrowdSaleCaller creates a new read-only instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSaleCaller(address common.Address, caller bind.ContractCaller) (*CrowdSaleCaller, error) {
	contract, err := bindCrowdSale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleCaller{contract: contract}, nil
}

// NewCrowdSaleTransactor creates a new write-only instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSaleTransactor(address common.Address, transactor bind.ContractTransactor) (*CrowdSaleTransactor, error) {
	contract, err := bindCrowdSale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleTransactor{contract: contract}, nil
}

// NewCrowdSaleFilterer creates a new log filterer instance of CrowdSale, bound to a specific deployed contract.
func NewCrowdSaleFilterer(address common.Address, filterer bind.ContractFilterer) (*CrowdSaleFilterer, error) {
	contract, err := bindCrowdSale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleFilterer{contract: contract}, nil
}

// bindCrowdSale binds a generic wrapper to an already deployed contract.
func bindCrowdSale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrowdSaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdSale *CrowdSaleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdSale.Contract.CrowdSaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdSale *CrowdSaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.Contract.CrowdSaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdSale *CrowdSaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdSale.Contract.CrowdSaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrowdSale *CrowdSaleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrowdSale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrowdSale *CrowdSaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrowdSale *CrowdSaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrowdSale.Contract.contract.Transact(opts, method, params...)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) EndTime() (*big.Int, error) {
	return _CrowdSale.Contract.EndTime(&_CrowdSale.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) EndTime() (*big.Int, error) {
	return _CrowdSale.Contract.EndTime(&_CrowdSale.CallOpts)
}

// FundsRaisedTotal is a free data retrieval call binding the contract method 0xeeeb3650.
//
// Solidity: function fundsRaisedTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) FundsRaisedTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "fundsRaisedTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsRaisedTotal is a free data retrieval call binding the contract method 0xeeeb3650.
//
// Solidity: function fundsRaisedTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) FundsRaisedTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsRaisedTotal(&_CrowdSale.CallOpts)
}

// FundsRaisedTotal is a free data retrieval call binding the contract method 0xeeeb3650.
//
// Solidity: function fundsRaisedTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) FundsRaisedTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsRaisedTotal(&_CrowdSale.CallOpts)
}

// FundsSoldTotal is a free data retrieval call binding the contract method 0xdb8c69ef.
//
// Solidity: function fundsSoldTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) FundsSoldTotal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "fundsSoldTotal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FundsSoldTotal is a free data retrieval call binding the contract method 0xdb8c69ef.
//
// Solidity: function fundsSoldTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) FundsSoldTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsSoldTotal(&_CrowdSale.CallOpts)
}

// FundsSoldTotal is a free data retrieval call binding the contract method 0xdb8c69ef.
//
// Solidity: function fundsSoldTotal() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) FundsSoldTotal() (*big.Int, error) {
	return _CrowdSale.Contract.FundsSoldTotal(&_CrowdSale.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsActive() (bool, error) {
	return _CrowdSale.Contract.IsActive(&_CrowdSale.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsActive() (bool, error) {
	return _CrowdSale.Contract.IsActive(&_CrowdSale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsFinished(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isFinished")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsFinished() (bool, error) {
	return _CrowdSale.Contract.IsFinished(&_CrowdSale.CallOpts)
}

// IsFinished is a free data retrieval call binding the contract method 0x7b352962.
//
// Solidity: function isFinished() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsFinished() (bool, error) {
	return _CrowdSale.Contract.IsFinished(&_CrowdSale.CallOpts)
}

// IsSalesCompleted is a free data retrieval call binding the contract method 0xa06c944a.
//
// Solidity: function isSalesCompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsSalesCompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isSalesCompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSalesCompleted is a free data retrieval call binding the contract method 0xa06c944a.
//
// Solidity: function isSalesCompleted() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsSalesCompleted() (bool, error) {
	return _CrowdSale.Contract.IsSalesCompleted(&_CrowdSale.CallOpts)
}

// IsSalesCompleted is a free data retrieval call binding the contract method 0xa06c944a.
//
// Solidity: function isSalesCompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsSalesCompleted() (bool, error) {
	return _CrowdSale.Contract.IsSalesCompleted(&_CrowdSale.CallOpts)
}

// IsUncompleted is a free data retrieval call binding the contract method 0xe2481cce.
//
// Solidity: function isUncompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCaller) IsUncompleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "isUncompleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUncompleted is a free data retrieval call binding the contract method 0xe2481cce.
//
// Solidity: function isUncompleted() view returns(bool)
func (_CrowdSale *CrowdSaleSession) IsUncompleted() (bool, error) {
	return _CrowdSale.Contract.IsUncompleted(&_CrowdSale.CallOpts)
}

// IsUncompleted is a free data retrieval call binding the contract method 0xe2481cce.
//
// Solidity: function isUncompleted() view returns(bool)
func (_CrowdSale *CrowdSaleCallerSession) IsUncompleted() (bool, error) {
	return _CrowdSale.Contract.IsUncompleted(&_CrowdSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrowdSale *CrowdSaleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrowdSale *CrowdSaleSession) Owner() (common.Address, error) {
	return _CrowdSale.Contract.Owner(&_CrowdSale.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CrowdSale *CrowdSaleCallerSession) Owner() (common.Address, error) {
	return _CrowdSale.Contract.Owner(&_CrowdSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) Price(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "price")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) Price() (*big.Int, error) {
	return _CrowdSale.Contract.Price(&_CrowdSale.CallOpts)
}

// Price is a free data retrieval call binding the contract method 0xa035b1fe.
//
// Solidity: function price() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) Price() (*big.Int, error) {
	return _CrowdSale.Contract.Price(&_CrowdSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) StartTime() (*big.Int, error) {
	return _CrowdSale.Contract.StartTime(&_CrowdSale.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) StartTime() (*big.Int, error) {
	return _CrowdSale.Contract.StartTime(&_CrowdSale.CallOpts)
}

// TargetFunds is a free data retrieval call binding the contract method 0x1459e2ad.
//
// Solidity: function targetFunds() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) TargetFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "targetFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetFunds is a free data retrieval call binding the contract method 0x1459e2ad.
//
// Solidity: function targetFunds() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) TargetFunds() (*big.Int, error) {
	return _CrowdSale.Contract.TargetFunds(&_CrowdSale.CallOpts)
}

// TargetFunds is a free data retrieval call binding the contract method 0x1459e2ad.
//
// Solidity: function targetFunds() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) TargetFunds() (*big.Int, error) {
	return _CrowdSale.Contract.TargetFunds(&_CrowdSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdSale *CrowdSaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdSale *CrowdSaleSession) Token() (common.Address, error) {
	return _CrowdSale.Contract.Token(&_CrowdSale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_CrowdSale *CrowdSaleCallerSession) Token() (common.Address, error) {
	return _CrowdSale.Contract.Token(&_CrowdSale.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_CrowdSale *CrowdSaleCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrowdSale.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_CrowdSale *CrowdSaleSession) TotalAmount() (*big.Int, error) {
	return _CrowdSale.Contract.TotalAmount(&_CrowdSale.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_CrowdSale *CrowdSaleCallerSession) TotalAmount() (*big.Int, error) {
	return _CrowdSale.Contract.TotalAmount(&_CrowdSale.CallOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_CrowdSale *CrowdSaleTransactor) Buy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "buy")
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_CrowdSale *CrowdSaleSession) Buy() (*types.Transaction, error) {
	return _CrowdSale.Contract.Buy(&_CrowdSale.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xa6f2ae3a.
//
// Solidity: function buy() payable returns()
func (_CrowdSale *CrowdSaleTransactorSession) Buy() (*types.Transaction, error) {
	return _CrowdSale.Contract.Buy(&_CrowdSale.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns(bool)
func (_CrowdSale *CrowdSaleTransactor) Refund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "refund")
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns(bool)
func (_CrowdSale *CrowdSaleSession) Refund() (*types.Transaction, error) {
	return _CrowdSale.Contract.Refund(&_CrowdSale.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x590e1ae3.
//
// Solidity: function refund() returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) Refund() (*types.Transaction, error) {
	return _CrowdSale.Contract.Refund(&_CrowdSale.TransactOpts)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xccb98ffc.
//
// Solidity: function setEndTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetEndTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setEndTime", _time)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xccb98ffc.
//
// Solidity: function setEndTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetEndTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetEndTime(&_CrowdSale.TransactOpts, _time)
}

// SetEndTime is a paid mutator transaction binding the contract method 0xccb98ffc.
//
// Solidity: function setEndTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetEndTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetEndTime(&_CrowdSale.TransactOpts, _time)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setPrice", _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetPrice(&_CrowdSale.TransactOpts, _price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x91b7f5ed.
//
// Solidity: function setPrice(uint256 _price) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetPrice(_price *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetPrice(&_CrowdSale.TransactOpts, _price)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactor) SetStartTime(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "setStartTime", _time)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleSession) SetStartTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetStartTime(&_CrowdSale.TransactOpts, _time)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _time) returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) SetStartTime(_time *big.Int) (*types.Transaction, error) {
	return _CrowdSale.Contract.SetStartTime(&_CrowdSale.TransactOpts, _time)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address addr) returns()
func (_CrowdSale *CrowdSaleTransactor) TransferOwnership(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "transferOwnership", addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address addr) returns()
func (_CrowdSale *CrowdSaleSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CrowdSale.Contract.TransferOwnership(&_CrowdSale.TransactOpts, addr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address addr) returns()
func (_CrowdSale *CrowdSaleTransactorSession) TransferOwnership(addr common.Address) (*types.Transaction, error) {
	return _CrowdSale.Contract.TransferOwnership(&_CrowdSale.TransactOpts, addr)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_CrowdSale *CrowdSaleTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_CrowdSale *CrowdSaleSession) Withdraw() (*types.Transaction, error) {
	return _CrowdSale.Contract.Withdraw(&_CrowdSale.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_CrowdSale *CrowdSaleTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CrowdSale.Contract.Withdraw(&_CrowdSale.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_CrowdSale *CrowdSaleTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CrowdSale.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_CrowdSale *CrowdSaleSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrowdSale.Contract.Fallback(&_CrowdSale.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_CrowdSale *CrowdSaleTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrowdSale.Contract.Fallback(&_CrowdSale.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdSale *CrowdSaleTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrowdSale.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdSale *CrowdSaleSession) Receive() (*types.Transaction, error) {
	return _CrowdSale.Contract.Receive(&_CrowdSale.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_CrowdSale *CrowdSaleTransactorSession) Receive() (*types.Transaction, error) {
	return _CrowdSale.Contract.Receive(&_CrowdSale.TransactOpts)
}

// CrowdSaleBoughtIterator is returned from FilterBought and is used to iterate over the raw logs and unpacked data for Bought events raised by the CrowdSale contract.
type CrowdSaleBoughtIterator struct {
	Event *CrowdSaleBought // Event containing the contract specifics and raw log

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
func (it *CrowdSaleBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdSaleBought)
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
		it.Event = new(CrowdSaleBought)
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
func (it *CrowdSaleBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdSaleBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdSaleBought represents a Bought event raised by the CrowdSale contract.
type CrowdSaleBought struct {
	Owner  common.Address
	Amount *big.Int
	Raise  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBought is a free log retrieval operation binding the contract event 0xa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c.
//
// Solidity: event Bought(address indexed owner, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) FilterBought(opts *bind.FilterOpts, owner []common.Address) (*CrowdSaleBoughtIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CrowdSale.contract.FilterLogs(opts, "Bought", ownerRule)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleBoughtIterator{contract: _CrowdSale.contract, event: "Bought", logs: logs, sub: sub}, nil
}

// WatchBought is a free log subscription operation binding the contract event 0xa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c.
//
// Solidity: event Bought(address indexed owner, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) WatchBought(opts *bind.WatchOpts, sink chan<- *CrowdSaleBought, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _CrowdSale.contract.WatchLogs(opts, "Bought", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdSaleBought)
				if err := _CrowdSale.contract.UnpackLog(event, "Bought", log); err != nil {
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

// ParseBought is a log parse operation binding the contract event 0xa9a40dec7a304e5915d11358b968c1e8d365992abf20f82285d1df1b30c8e24c.
//
// Solidity: event Bought(address indexed owner, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) ParseBought(log types.Log) (*CrowdSaleBought, error) {
	event := new(CrowdSaleBought)
	if err := _CrowdSale.contract.UnpackLog(event, "Bought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdSaleRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the CrowdSale contract.
type CrowdSaleRefundedIterator struct {
	Event *CrowdSaleRefunded // Event containing the contract specifics and raw log

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
func (it *CrowdSaleRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdSaleRefunded)
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
		it.Event = new(CrowdSaleRefunded)
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
func (it *CrowdSaleRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdSaleRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdSaleRefunded represents a Refunded event raised by the CrowdSale contract.
type CrowdSaleRefunded struct {
	From   common.Address
	Amount *big.Int
	Raise  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66.
//
// Solidity: event Refunded(address indexed from, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) FilterRefunded(opts *bind.FilterOpts, from []common.Address) (*CrowdSaleRefundedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.FilterLogs(opts, "Refunded", fromRule)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleRefundedIterator{contract: _CrowdSale.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66.
//
// Solidity: event Refunded(address indexed from, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *CrowdSaleRefunded, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.WatchLogs(opts, "Refunded", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdSaleRefunded)
				if err := _CrowdSale.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x2dc8e290002f06fc0085bbca9dfb8b415cf4d1178950c72ff9ee8f4d8878ee66.
//
// Solidity: event Refunded(address indexed from, uint256 amount, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) ParseRefunded(log types.Log) (*CrowdSaleRefunded, error) {
	event := new(CrowdSaleRefunded)
	if err := _CrowdSale.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrowdSaleWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the CrowdSale contract.
type CrowdSaleWithdrawnIterator struct {
	Event *CrowdSaleWithdrawn // Event containing the contract specifics and raw log

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
func (it *CrowdSaleWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrowdSaleWithdrawn)
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
		it.Event = new(CrowdSaleWithdrawn)
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
func (it *CrowdSaleWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrowdSaleWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrowdSaleWithdrawn represents a Withdrawn event raised by the CrowdSale contract.
type CrowdSaleWithdrawn struct {
	From  common.Address
	Raise *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed from, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) FilterWithdrawn(opts *bind.FilterOpts, from []common.Address) (*CrowdSaleWithdrawnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.FilterLogs(opts, "Withdrawn", fromRule)
	if err != nil {
		return nil, err
	}
	return &CrowdSaleWithdrawnIterator{contract: _CrowdSale.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed from, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *CrowdSaleWithdrawn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _CrowdSale.contract.WatchLogs(opts, "Withdrawn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrowdSaleWithdrawn)
				if err := _CrowdSale.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed from, uint256 raise)
func (_CrowdSale *CrowdSaleFilterer) ParseWithdrawn(log types.Log) (*CrowdSaleWithdrawn, error) {
	event := new(CrowdSaleWithdrawn)
	if err := _CrowdSale.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}