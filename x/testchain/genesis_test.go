package testchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/yogeshguptakaithal/testchain/testutil/keeper"
	"github.com/yogeshguptakaithal/testchain/testutil/nullify"
	"github.com/yogeshguptakaithal/testchain/x/testchain"
	"github.com/yogeshguptakaithal/testchain/x/testchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		CommentList: []types.Comment{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		CommentCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TestchainKeeper(t)
	testchain.InitGenesis(ctx, *k, genesisState)
	got := testchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.CommentList, got.CommentList)
	require.Equal(t, genesisState.CommentCount, got.CommentCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
