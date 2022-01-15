package keeper

import (
	"github.com/cosmonaut/interchange/x/dex/types"
)

var _ types.QueryServer = Keeper{}
