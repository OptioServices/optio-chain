package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/OptioServices/optio-chain/testutil/keeper"
	"github.com/OptioServices/optio-chain/x/distribute/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.DistributeKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
