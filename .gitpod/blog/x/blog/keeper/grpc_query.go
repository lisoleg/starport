package keeper

import (
	"github.com/lisoleg/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
