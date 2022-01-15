package hello_test

import (
	"testing"

	keepertest "github.com/cosmonaut/hello/testutil/keeper"
	"github.com/cosmonaut/hello/x/hello"
	"github.com/cosmonaut/hello/x/hello/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloKeeper(t)
	hello.InitGenesis(ctx, *k, genesisState)
	got := hello.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
