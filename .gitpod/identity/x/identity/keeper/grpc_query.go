package keeper

import (
	"github.com/lisoleg/identity/x/identity/types"
)

var _ types.QueryServer = Keeper{}
