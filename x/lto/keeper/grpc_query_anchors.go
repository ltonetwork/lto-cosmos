package keeper

import (
	"context"

	"lto/x/lto/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Anchors(goCtx context.Context, req *types.QueryAnchorsRequest) (*types.QueryAnchorsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	// define a variable that will store a list of anchors
	var anchors []*types.Anchor

	// get context with the information about the env
	ctx := sdk.UnwrapSDKContext(goCtx)

	// get the kv module store using the store key
	store := ctx.KVStore(k.storeKey)

	// get the part of the store that keeps anchors
	anchorStore := prefix.NewStore(store, []byte(types.AnchorKey))

	// paginate the anchors store based on PageRequest
	pageRes, err := query.Paginate(anchorStore, req.Pagination, func(key []byte, value []byte) error {
		var anchor types.Anchor
		if err := k.cdc.Unmarshal(value, &anchor); err != nil {
			return err
		}

		anchors = append(anchors, &anchor)

		return nil
	})

	// throw an error if pagination failed
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// return a struct containing a list of anchors and pagination info
	return &types.QueryAnchorsResponse{Anchor: anchors, Pagination: pageRes}, nil
}
