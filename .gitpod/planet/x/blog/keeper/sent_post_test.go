package keeper_test

import (
	"testing"

	keepertest "github.com/cosmonaut/planet/testutil/keeper"
	"github.com/cosmonaut/planet/x/blog/keeper"
	"github.com/cosmonaut/planet/x/blog/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNSentPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentPost {
	items := make([]types.SentPost, n)
	for i := range items {
		items[i].Id = keeper.AppendSentPost(ctx, items[i])
	}
	return items
}

func TestSentPostGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentPost(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentPost(ctx, item.Id)
		require.True(t, found)
		require.Equal(t, item, got)
	}
}

func TestSentPostRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentPost(ctx, item.Id)
		_, found := keeper.GetSentPost(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentPost(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllSentPost(ctx))
}

func TestSentPostCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSentPost(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentPostCount(ctx))
}
