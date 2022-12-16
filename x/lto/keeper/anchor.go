package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"lto/x/lto/types"
)

func (k Keeper) AppendAnchor(ctx sdk.Context, anchor types.Anchor, hash string) bool {

	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AnchorKey))

	anchor.Timestamp = uint32(ctx.BlockHeight())

	hashBytes := []byte(hash)
	/*
		bytes txType = 1;
		bytes version = 2;
		bytes chainId = 3;
		uint32 timestamp = 4;
		bytes keyType = 5;
		bytes pubKey = 6;
		uint32 fee = 7;
		uint32 anchorCount = 8;
		bytes anchorLength = 9;
		bytes anchors = 10;
		}
	*/
	appendedValue := k.cdc.MustMarshal(&anchor)

	store.Set(hashBytes, appendedValue)

	return true
}
