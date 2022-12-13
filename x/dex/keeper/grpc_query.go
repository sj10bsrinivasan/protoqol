package keeper

import (
	"protoqol/x/dex/types"
)

var _ types.QueryServer = Keeper{}
