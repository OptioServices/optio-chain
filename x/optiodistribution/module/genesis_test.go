package optiodistribution_test

import (
	"testing"

	keepertest "optio/testutil/keeper"
	"optio/testutil/nullify"
	optiodistribution "optio/x/optiodistribution/module"
	"optio/x/optiodistribution/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.OptiodistributionKeeper(t)
	optiodistribution.InitGenesis(ctx, k, genesisState)
	got := optiodistribution.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
