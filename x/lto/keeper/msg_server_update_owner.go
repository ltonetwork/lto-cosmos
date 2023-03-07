package keeper

import (
	"context"

	"lto-cosmos/x/lto/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)

	// check for the existance of denom
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
	}

	// check for owner
	if valFound.Owner != msg.Owner {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "only owner can update")
	}

	// update the appropriate fields of the denom
	var denom = types.Denom{
		Owner:              msg.NewOwner,
		Denom:              msg.Denom,
		Description:        valFound.Description,
		MaxSupply:          valFound.MaxSupply,
		Supply:             valFound.Supply,
		Precision:          valFound.Precision,
		Ticker:             valFound.Ticker,
		Url:                valFound.Url,
		CanChangeMaxSupply: valFound.CanChangeMaxSupply,
	}

	// save denom in the keeper
	k.SetDenom(
		ctx,
		denom,
	)

	return &types.MsgUpdateOwnerResponse{}, nil
}
