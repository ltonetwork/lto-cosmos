package keeper

import (
	"context"

	"lto/x/lto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAnchor(goCtx context.Context, msg *types.MsgCreateAnchor) (*types.MsgCreateAnchorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var anchor = types.Anchor{
		Creator:  msg.Creator,
		Hash:     msg.Hash,
		Encoding: msg.Encoding,
	}

	// add an anchor to the store and get its ID
	id := k.AppendAnchor(ctx, anchor)

	// return the ID
	return &types.MsgCreateAnchorResponse{Id: id}, nil
}
