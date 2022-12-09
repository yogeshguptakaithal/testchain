package testchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/yogeshguptakaithal/testchain/testutil/sample"
	testchainsimulation "github.com/yogeshguptakaithal/testchain/x/testchain/simulation"
	"github.com/yogeshguptakaithal/testchain/x/testchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = testchainsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgCreateComment = "op_weight_msg_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateComment int = 100

	opWeightMsgUpdateComment = "op_weight_msg_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateComment int = 100

	opWeightMsgDeleteComment = "op_weight_msg_comment"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteComment int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	testchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		CommentList: []types.Comment{
			{
				Id:      0,
				Creator: sample.AccAddress(),
			},
			{
				Id:      1,
				Creator: sample.AccAddress(),
			},
		},
		CommentCount: 2,
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&testchainGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateComment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateComment, &weightMsgCreateComment, nil,
		func(_ *rand.Rand) {
			weightMsgCreateComment = defaultWeightMsgCreateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateComment,
		testchainsimulation.SimulateMsgCreateComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateComment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateComment, &weightMsgUpdateComment, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateComment = defaultWeightMsgUpdateComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateComment,
		testchainsimulation.SimulateMsgUpdateComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteComment int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteComment, &weightMsgDeleteComment, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteComment = defaultWeightMsgDeleteComment
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteComment,
		testchainsimulation.SimulateMsgDeleteComment(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
