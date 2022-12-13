package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "protoqol/testutil/keeper"
	"protoqol/x/protoqol/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ProtoqolKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
