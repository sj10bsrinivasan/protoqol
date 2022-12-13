package protoqol_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "protoqol/testutil/keeper"
	"protoqol/testutil/nullify"
	"protoqol/x/protoqol"
	"protoqol/x/protoqol/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProtoqolKeeper(t)
	protoqol.InitGenesis(ctx, *k, genesisState)
	got := protoqol.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
