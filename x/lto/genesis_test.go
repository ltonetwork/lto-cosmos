package lto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lto-cosmos/testutil/keeper"
	"lto-cosmos/testutil/nullify"
	"lto-cosmos/x/lto"
	"lto-cosmos/x/lto/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DenomList: []types.Denom{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LtoKeeper(t)
	lto.InitGenesis(ctx, *k, genesisState)
	got := lto.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
