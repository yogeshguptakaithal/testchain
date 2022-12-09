package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/yogeshguptakaithal/testchain/testutil/keeper"
	"github.com/yogeshguptakaithal/testchain/x/testchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.TestchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
