package api

import (
	"log"
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

func (handler *ERC20TokenAPIHandler) ERC20TokenNew(opts model.ERC20TokenOpts, contract *model.Contract) error  {
	log.Println("server\t-", "Receive GetName call, id:", opts.Name)
	contract.Address = "abc"
	return nil
	//jsonapiRuntime := jsonapi.NewRuntime().Instrument("blogs.list")

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

