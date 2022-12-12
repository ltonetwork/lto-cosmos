package keeper

import (
	"context"

	"lto/x/lto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) NetworkType(goCtx context.Context, req *types.QueryNetworkTypeRequest) (*types.QueryNetworkTypeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	return &types.QueryNetworkTypeResponse{Text: "T"}, nil
}
