package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/runtime"

	"github.com/OptioServices/optio/x/distribute/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx context.Context) (params types.Params) {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz := store.Get(types.ParamsKey)
	if bz == nil {
		return params
	}

	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// SetParams set the params
func (k Keeper) SetParams(ctx context.Context, params types.Params) error {
	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	bz, err := k.cdc.Marshal(&params)
	if err != nil {
		return err
	}
	store.Set(types.ParamsKey, bz)

	return nil
}

func (k Keeper) IsAuthorized(ctx context.Context, address string) bool {
	authorizedAccounts := k.GetParams(ctx).AuthorizedAccounts
	for _, authorizedAccount := range authorizedAccounts {
		if authorizedAccount == address {
			return true
		}
	}

	return false
}
