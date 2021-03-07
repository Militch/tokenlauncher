package tokenlauncher

import (
	"github.com/ethereum/go-ethereum/common"
	"tokenlauncher/uint256"
)

type CrowdSale struct {
	TokenAddress common.Address
	TargetFunds uint256.UInt256
	Price uint256.UInt256
	
}