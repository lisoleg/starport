package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cosmonaut/interchange/testutil/keeper"
	"github.com/cosmonaut/interchange/x/dex/keeper"
	"github.com/cosmonaut/interchange/x/dex/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.DexKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
