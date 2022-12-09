package lto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lto/testutil/keeper"
	"lto/testutil/nullify"
	"lto/x/lto"
	"lto/x/lto/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LtoKeeper(t)
	lto.InitGenesis(ctx, *k, genesisState)
	got := lto.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}