package model

import (
	"github.com/ethereum/go-ethereum/common"
	"tokenlauncher/uint256"
)

type CrowdSaleOpts struct {
	TokenAddress common.Address
	TargetFunds uint256.UInt256
	Price uint256.UInt256
	StartTime uint256.UInt256
	EndTime uint256.UInt256
}