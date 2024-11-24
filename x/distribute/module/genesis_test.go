package distribute_test

import (
	"testing"

	keepertest "github.com/OptioServices/optio-chain/testutil/keeper"
	"github.com/OptioServices/optio-chain/testutil/nullify"
	distribute "github.com/OptioServices/optio-chain/x/distribute/module"
	"github.com/OptioServices/optio-chain/x/distribute/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.DistributeKeeper(t)
	distribute.InitGenesis(ctx, k, genesisState)
	got := distribute.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
