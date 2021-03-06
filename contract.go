package tokenlauncher

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type Contract struct {
	Data string
	ABI string
	Address string
}




func (c *Contract) BuildTransaction() *types.Transaction {

	return types.NewTx(&types.LegacyTx{

	})
}