package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "lto-cosmos/testutil/keeper"
	"lto-cosmos/x/lto/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.LtoKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
