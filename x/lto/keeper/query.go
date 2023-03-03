package keeper

import (
	"lto-cosmos/x/lto/types"
)

var _ types.QueryServer = Keeper{}
