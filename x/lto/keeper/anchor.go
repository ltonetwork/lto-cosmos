package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"lto/x/lto/types"
)

func (k Keeper) AppendAnchor(ctx sdk.Context, anchor types.Anchor) uint64 {
	// get the current number of anchors
	count := k.GetAnchorCount(ctx)

	// assign an ID
	anchor.Id = count

	// get the store
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AnchorKey))

	// convert the anchor ID into bytes
	byteKey := make([]byte, 8)
	binary.BigEndian.PutUint64(byteKey, anchor.Id)

	// marshall the anchor into bytes
	appendedValue := k.cdc.MustMarshal(&anchor)

	// insert the anchor bytes using anchor ID as key
	store.Set(byteKey, appendedValue)

	// update the anchor count
	k.SetAnchorCount(ctx, count+1)
	return count
}

func (k Keeper) GetAnchorCount(ctx sdk.Context) uint64 {
	// get the store using storeKey ("lto") and AnchorCountKey ("Anchor/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AnchorCountKey))

	// convert AnchorCountKey to bytes
	byteKey := []byte(types.AnchorCountKey)

	// get the count value
	bz := store.Get(byteKey)

	// return zero if the count value is not found
	if bz == nil {
		return 0
	}

	// convert the count into a uint64
	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetAnchorCount(ctx sdk.Context, count uint64) {
	// get the store using storeKey ("lto") and AnchorCountKey ("Anchor/count/")
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.AnchorCountKey))

	// convert the AnchorCountKey to bytes
	byteKey := []byte(types.AnchorCountKey)

	// convert count from uint64 to string and get bytes
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)

	// set the value of Anchor/count/ to count
	store.Set(byteKey, bz)
}
