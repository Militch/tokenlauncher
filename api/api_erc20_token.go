package api

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"tokenlauncher/launcher"
	"tokenlauncher/uint256"
)

type ERC20TokenAPIHandler struct {

}



type ERC20TokenReq struct {
	Name string
	Symbol string
	Decimals string
	InitialSupply string
}

type ContractResp struct {
	Address string `json:"address"`
	TxHash string `json:"txHash"`
	ABI string `json:"abi"`
	Data string `json:"data"`
}

type ERC20TokenValueGet struct {
	Address string
}

type ERC20TokenOwnerGet struct {
	Address string
	Owner string
}

type ERC20TokenAllowanceGet struct {
	Address string
	Owner string
	Spender string
}


type ERC20TokenApproveSet struct {
	Address string
	Spender string
	Value string
}

type ERC20TokenBurnSet struct {
	Address string
	Value string
}

type ERC20TokenBurnFromSet struct {
	Address string
	From string
	Value string
}

type ERC20TokenTransferSet struct {
	Address string
	To string
	Value string
}

type ERC20TokenTransferFromSet struct {
	Address string
	From string
	To string
	Value string
}

type ERC20TokenValueGot struct {
	Value string `json:"value"`
}



type ERC20TokenTxGot struct {
	Hash string `json:"hash"`
}



func (handler *ERC20TokenAPIHandler) NewAndDeploy(tk ERC20TokenReq, resp *ContractResp) error  {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	decimals,err := strconv.ParseInt(tk.Decimals,10,8)
	tkOpts := &launcher.ERC20TokenOpts{
		Name: tk.Name,
		Symbol: tk.Symbol,
		InitialSupply: tk.InitialSupply,
		Decimals:uint8(decimals),
	}
	contract, err := la.DeployERC20Token(tkOpts)
	if err != nil {
		return err
	}
	resp.Address = contract.Address
	resp.TxHash = contract.TxHash
	resp.ABI = contract.ABI
	resp.Data = contract.Data
	return nil
}

func (handler *ERC20TokenAPIHandler) GetName(get ERC20TokenValueGet, got *ERC20TokenValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := erc20token.Name(nil)
	if err != nil {
		return err
	}
	got.Value = value
	return nil
}

func (handler *ERC20TokenAPIHandler) GetSymbol(get ERC20TokenValueGet, got *ERC20TokenValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := erc20token.Symbol(nil)
	if err != nil {
		return err
	}
	got.Value = value
	return nil
}

func (handler *ERC20TokenAPIHandler) GetDecimals(get ERC20TokenValueGet, got *ERC20TokenValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := erc20token.Decimals(nil)
	if err != nil {
		return err
	}
	got.Value = strconv.Itoa(int(value))
	return nil
}

func (handler *ERC20TokenAPIHandler) GetTotalSupply(get ERC20TokenValueGet, got *ERC20TokenValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := erc20token.TotalSupply(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func (handler *ERC20TokenAPIHandler) BalanceOf(get ERC20TokenOwnerGet, got *ERC20TokenValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	ownerAddr := common.HexToAddress(get.Owner)
	value,err := erc20token.BalanceOf(nil, ownerAddr )
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func (handler *ERC20TokenAPIHandler) GetAllowance(get ERC20TokenAllowanceGet, got *ERC20TokenValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	ownerAddr := common.HexToAddress(get.Owner)
	spenderAddr := common.HexToAddress(get.Spender)
	value,err := erc20token.Allowance(nil, ownerAddr, spenderAddr)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func (handler *ERC20TokenAPIHandler) Approve(set ERC20TokenApproveSet, got *ERC20TokenTxGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	spenderAddr := common.HexToAddress(set.Spender)
	value := uint256.NewUInt256(set.Value)
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	tx, err := erc20token.Approve(txOpts,spenderAddr, value.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

func (handler *ERC20TokenAPIHandler) Burn(set ERC20TokenBurnSet, got *ERC20TokenTxGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value := uint256.NewUInt256(set.Value)
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	tx, err := erc20token.Burn(txOpts, value.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

func (handler *ERC20TokenAPIHandler) BurnFrom(set ERC20TokenBurnFromSet, got *ERC20TokenTxGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	fromAddr := common.HexToAddress(set.From)
	value := uint256.NewUInt256(set.Value)
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	tx, err := erc20token.BurnFrom(txOpts, fromAddr, value.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

func (handler *ERC20TokenAPIHandler) Transfer(set ERC20TokenTransferSet, got *ERC20TokenTxGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	toAddr := common.HexToAddress(set.To)
	value := uint256.NewUInt256(set.Value)
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	tx, err := erc20token.Transfer(txOpts, toAddr, value.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

func (handler *ERC20TokenAPIHandler) TransferFrom(set ERC20TokenTransferFromSet, got *ERC20TokenTxGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	erc20token, err := la.LoadERC20TokenByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	fromAddr := common.HexToAddress(set.From)
	toAddr := common.HexToAddress(set.To)
	value := uint256.NewUInt256(set.Value)
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	tx, err := erc20token.TransferFrom(txOpts, fromAddr, toAddr, value.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

