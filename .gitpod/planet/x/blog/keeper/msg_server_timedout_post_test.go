package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/cosmonaut/planet/x/blog/types"
)

func TestTimedoutPostMsgServerCreate(t *testing.T) {
	srv, ctx := setupMsgServer(t)
	creator := "A"
	for i := 0; i < 5; i++ {
		resp, err := srv.CreateTimedoutPost(ctx, &types.MsgCreateTimedoutPost{Creator: creator})
		require.NoError(t, err)
		require.Equal(t, i, int(resp.Id))
	}
}

func TestTimedoutPostMsgServerUpdate(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgUpdateTimedoutPost
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgUpdateTimedoutPost{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutPost{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgUpdateTimedoutPost{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)
			_, err := srv.CreateTimedoutPost(ctx, &types.MsgCreateTimedoutPost{Creator: creator})
			require.NoError(t, err)

			_, err = srv.UpdateTimedoutPost(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestTimedoutPostMsgServerDelete(t *testing.T) {
	creator := "A"

	for _, tc := range []struct {
		desc    string
		request *types.MsgDeleteTimedoutPost
		err     error
	}{
		{
			desc:    "Completed",
			request: &types.MsgDeleteTimedoutPost{Creator: creator},
		},
		{
			desc:    "Unauthorized",
			request: &types.MsgDeleteTimedoutPost{Creator: "B"},
			err:     sdkerrors.ErrUnauthorized,
		},
		{
			desc:    "KeyNotFound",
			request: &types.MsgDeleteTimedoutPost{Creator: creator, Id: 10},
			err:     sdkerrors.ErrKeyNotFound,
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			srv, ctx := setupMsgServer(t)

			_, err := srv.CreateTimedoutPost(ctx, &types.MsgCreateTimedoutPost{Creator: creator})
			require.NoError(t, err)
			_, err = srv.DeleteTimedoutPost(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
