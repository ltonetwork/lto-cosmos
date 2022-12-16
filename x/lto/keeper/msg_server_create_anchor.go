package keeper

import (
	"context"

	"lto/x/lto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateAnchor(goCtx context.Context, msg *types.MsgCreateAnchor) (*types.MsgCreateAnchorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	print("msg.getHash(): ", msg.GetHash())
	/*
		byte txType = 1;
		byte version = 2;
		byte chainId = 3;
		uint32 timestamp = 4;
		byte keyType = 5;
		[]byte pubKey = 6;
		uint32 fee = 7;
		uint16 anchorCount = 8;
		[]byte anchorLength = 9;
		[]byte anchors = 10;
	*/
	var anchor = types.Anchor{
		PubKey:   msg.Creator,
		Anchor:   msg.Hash,
		Encoding: msg.Encoding,
	}

	// add an anchor to the store and get its ID
	success := k.AppendAnchor(ctx, anchor, msg.GetHash())

	// return the status
	return &types.MsgCreateAnchorResponse{Success: success}, nil
}
