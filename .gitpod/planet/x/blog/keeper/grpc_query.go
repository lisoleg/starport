package keeper

import (
	"github.com/cosmonaut/planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
