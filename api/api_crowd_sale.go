package api

import (
	"github.com/ethereum/go-ethereum/common"
	"strconv"
	"tokenlauncher/launcher"
	"tokenlauncher/uint256"
)

type CrowdSaleAPIHandler struct {

}


type CrowdSaleReq struct {
	TokenAddress string
	TargetFunds string
	Price string
	StartTime string
	EndTime string
}
type CrowdSaleValueGet struct {
	Address string
}



type CrowdSaleValueGot struct {
	Value string `json:"value"`
}
type CrowdSaleWithdrawSet struct {
	Address string
}

type CrowdSaleTxGot struct {
	Hash string `json:"hash"`
}

type CrowdSaleTimeSet struct {
	Address string
	Time string
}
type CrowdSalePriceSet struct {
	Address string
	Price string
}
type CrowdSaleTargetFundsSet struct {
	Address string
	Amount string
}

func (handler *CrowdSaleAPIHandler) NewAndDeploy(req CrowdSaleReq, resp *ContractResp) error  {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	tokenAddr := common.HexToAddress(req.TokenAddress)
	targetFunds := *uint256.NewUInt256(req.TargetFunds)
	price := *uint256.NewUInt256(req.Price)
	startTime := *uint256.NewUInt256(req.StartTime)
	endTime := *uint256.NewUInt256(req.EndTime)
	opts := &launcher.CrowdSaleOpts{
		TokenAddress: tokenAddr,
		TargetFunds: targetFunds,
		Price: price,
		StartTime: startTime,
		EndTime: endTime,
	}
	contract, err := la.DeployCrowdSale(opts)
	if err != nil {
		return err
	}
	resp.Address = contract.Address
	resp.TxHash = contract.TxHash
	resp.ABI = contract.ABI
	resp.Data = contract.Data
	return nil
}


func (handler *CrowdSaleAPIHandler) GetFundsSoldTotal(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {

	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.FundsSoldTotal(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func (handler *CrowdSaleAPIHandler) GetFundsRaisedTotal(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {

	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.FundsRaisedTotal(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func (handler *CrowdSaleAPIHandler) GetToken(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {

	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.Token(nil)
	if err != nil {
		return err
	}
	got.Value = value.Hex()
	return nil
}

func (handler *CrowdSaleAPIHandler) GetPrice(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {

	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.Price(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}


func (handler *CrowdSaleAPIHandler) GetTargetFunds(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.TargetFunds(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}


func (handler *CrowdSaleAPIHandler) GetStartTime(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.StartTime(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}


func (handler *CrowdSaleAPIHandler) GetEndTime(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
    la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.EndTime(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func (handler *CrowdSaleAPIHandler) GetTotalAmount(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.TotalAmount(nil)
	if err != nil {
		return err
	}
	got.Value = "0x" + value.Text(16)
	return nil
}

func btoi(b bool) int {
	i := 0
	if b  {
		i = 1
	}
	return i
}

func (handler *CrowdSaleAPIHandler) IsActive(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.IsActive(nil)
	if err != nil {
		return err
	}
	got.Value = strconv.Itoa(btoi(value))
	return nil
}

func (handler *CrowdSaleAPIHandler) IsFinished(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.IsFinished(nil)
	if err != nil {
		return err
	}
	got.Value = strconv.Itoa(btoi(value))
	return nil
}


func (handler *CrowdSaleAPIHandler) IsUncompleted(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.IsUncompleted(nil)
	if err != nil {
		return err
	}
	got.Value = strconv.Itoa(btoi(value))
	return nil
}


func (handler *CrowdSaleAPIHandler) IsSalesCompleted(get CrowdSaleValueGet, got *CrowdSaleValueGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := get.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	value,err := crowdSale.IsSalesCompleted(nil)
	if err != nil {
		return err
	}
	got.Value = strconv.Itoa(btoi(value))
	return nil
}

func (handler *CrowdSaleAPIHandler) Withdraw(set CrowdSaleWithdrawSet, got *CrowdSaleTxGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	tx, err := crowdSale.Withdraw(txOpts)
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}


func (handler *CrowdSaleAPIHandler) SetStartTime(set CrowdSaleTimeSet, got *CrowdSaleTxGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	time := uint256.NewUInt256(set.Time)
	tx, err := crowdSale.SetStartTime(txOpts,time.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

func (handler *CrowdSaleAPIHandler) SetEndTime(set CrowdSaleTimeSet, got *CrowdSaleTxGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	time := uint256.NewUInt256(set.Time)
	tx, err := crowdSale.SetEndTime(txOpts,time.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}

func (handler *CrowdSaleAPIHandler) SetPrice(set CrowdSalePriceSet, got *CrowdSaleTxGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	price := uint256.NewUInt256(set.Price)
	tx, err := crowdSale.SetPrice(txOpts,price.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}


func (handler *CrowdSaleAPIHandler) SetTargetFunds(set CrowdSaleTargetFundsSet, got *CrowdSaleTxGot) error {
	la, err := launcher.NewLauncherDefault()
	if err != nil {
		return err
	}
	contractAddr := set.Address
	crowdSale, err := la.LoadCrowdSaleByContractAddr(contractAddr)
	if err != nil {
		return err
	}
	txOpts, err := la.DefaultTxOpts()
	if err != nil {
		return err
	}
	amount := uint256.NewUInt256(set.Amount)
	tx, err := crowdSale.SetTargetFunds(txOpts,amount.ToBigInt())
	if err != nil {
		return err
	}
	got.Hash = tx.Hash().Hex()
	return nil
}