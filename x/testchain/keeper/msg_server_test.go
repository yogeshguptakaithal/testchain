package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/yogeshguptakaithal/testchain/testutil/keeper"
	"github.com/yogeshguptakaithal/testchain/x/testchain/keeper"
	"github.com/yogeshguptakaithal/testchain/x/testchain/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.TestchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
