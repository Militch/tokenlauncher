package tokenlauncher

import (
	"testing"
	"tokenlauncher/uint256"
)

func TestNewERC20TokenContract(t *testing.T) {
	totalSupply := *uint256.NewUInt256ByUInt32(uint32(2000000000))
	token := &ERC20Token{
		Name: "FixToken",
		Symbol: "FIX",
		InitialSupply: totalSupply,
		Decimals: uint8(12),
	}
	c, err := NewERC20TokenContract(token)
	if err != nil {
		t.Error(err)
	}
	_ = c
}