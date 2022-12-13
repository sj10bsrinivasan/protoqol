package keeper

import (
	"protoqol/x/blog/types"
)

var _ types.QueryServer = Keeper{}
