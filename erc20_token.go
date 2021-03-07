package tokenlauncher

import "tokenlauncher/uint256"

type ERC20Token struct {
	Name          string
	Symbol        string
	Decimals      uint8
	InitialSupply uint256.UInt256
	launcher      Launcher
}

func NewERC20TokenByContract(contract *Contract) *ERC20Token {

	return &ERC20Token{

	}
}

func (token *ERC20Token) Transfer(
	address uint256.UInt256,
	value uint256.UInt256)  {

}

func (token *ERC20Token) TransferFrom(
	from uint256.UInt256, to uint256.UInt256,
	value uint256.UInt256)  {

}
