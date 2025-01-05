package keeper

import (
	"legal/x/legal/types"
)

var _ types.QueryServer = Keeper{}
