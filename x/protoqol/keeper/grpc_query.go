package keeper

import (
	"protoqol/x/protoqol/types"
)

var _ types.QueryServer = Keeper{}
