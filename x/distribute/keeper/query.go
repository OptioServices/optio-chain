package keeper

import (
	"github.com/OptioServices/optio-chain/x/distribute/types"
)

var _ types.QueryServer = Keeper{}
