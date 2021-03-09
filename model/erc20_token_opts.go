package model

type ERC20TokenOpts struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	Decimals      uint8 `json:"decimals"`
	InitialSupply string `json:"initialSupply"`
}

