package api

import (
	"math/big"
	"strconv"
	"tokenlauncher/launcher"
	"tokenlauncher/model"
)

type ERC20TokenAPIHandler struct {

}

type ERC20TokenSender struct {
	Address string `json:"address"`
}

type ERC20TokenGet struct {
	Address string `json:"address"`
	Params string `json:"params"`
}

type ERC20TokenGot struct {
	Address string `json:"address"`
}

type ERC20Token struct {
	Name          string
	Symbol        string
	Decimals      string
	InitialSupply string
}



func (handler *ERC20TokenAPIHandler) NewAndDeploy(tk ERC20Token, contract *model.Contract) error  {
	rcpUrl := "http://localhost:8101"
	keyFile := "D:\\workspace\\fixtoken\\data0\\keystore\\UTC--2021-03-05T09-35-51.816544100Z--60d1148b3b2ab38a5937dc30244a3b4c5ec6da52"
	password := "123"
	la, err := launcher.NewLauncher(rcpUrl,big.NewInt(10086),keyFile,password)
	if err != nil {
		return err
	}
	decimals,err := strconv.ParseInt(tk.Decimals,10,8)
	tkOpts := &model.ERC20TokenOpts{
		Name: tk.Name,
		Symbol: tk.Symbol,
		InitialSupply: tk.InitialSupply,
		Decimals:uint8(decimals),
	}
	r, err := la.DeployERC20Token(tkOpts)
	if err != nil {
		return err
	}
	_ = r
	return nil
}

func (handler *ERC20TokenAPIHandler) GetName(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetName"
	return nil
}

func (handler *ERC20TokenAPIHandler) GetSymbol(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) GetDecimals(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) GetTotalSupply(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) BalanceOf(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}


func (handler *ERC20TokenAPIHandler) Approve(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) Burn(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) BurnFrom(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) Transfer(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

func (handler *ERC20TokenAPIHandler) TransferFrom(get ERC20TokenGet, got *ERC20TokenGot) error {
	got.Address = "GetSymbol"
	return nil
}

