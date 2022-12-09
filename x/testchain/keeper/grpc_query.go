package keeper

import (
	"github.com/yogeshguptakaithal/testchain/x/testchain/types"
)

var _ types.QueryServer = Keeper{}
