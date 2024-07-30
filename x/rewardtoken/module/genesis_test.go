package rewardtoken_test

import (
	"testing"

	keepertest "optio/testutil/keeper"
	"optio/testutil/nullify"
	rewardtoken "optio/x/rewardtoken/module"
	"optio/x/rewardtoken/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		DenomList: []types.Denom{
			{
				Denom: "0",
			},
			{
				Denom: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RewardtokenKeeper(t)
	rewardtoken.InitGenesis(ctx, k, genesisState)
	got := rewardtoken.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.DenomList, got.DenomList)
	// this line is used by starport scaffolding # genesis/test/assert
}
