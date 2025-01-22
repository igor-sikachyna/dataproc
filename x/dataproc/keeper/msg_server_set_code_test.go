package keeper_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	keepertest "github.com/igor-sikachyna/dataproc/testutil/keeper"
	dataproc "github.com/igor-sikachyna/dataproc/x/dataproc"
	"github.com/igor-sikachyna/dataproc/x/dataproc/keeper"
	"github.com/igor-sikachyna/dataproc/x/dataproc/testutil"
)

func setupMsgServerSetCode(t testing.TB) (dataproc.MsgServer, keeper.Keeper, context.Context) {
	server, k, context, _, _ := setupMsgServerSetCodeWithMock(t)
	return server, k, context
}

func setupMsgServerSetCodeWithMock(t testing.TB) (dataproc.MsgServer, keeper.Keeper, context.Context,
	*gomock.Controller, *testutil.MockBankKeeper) {
	ctrl := gomock.NewController(t)
	bankMock := testutil.NewMockBankKeeper(ctrl)
	k, ctx := keepertest.CheckersKeeperWithMocks(t)
	k.InitGenesis(ctx, dataproc.NewGenesisState())
	server := keeper.NewMsgServerImpl(k)
	return server, k, ctx, ctrl, bankMock
}

func TestSetCode(t *testing.T) {
	msgServer, _, context := setupMsgServerSetCode(t)
	createResponse, err := msgServer.SetCode(context, &dataproc.MsgSetCode{
		Creator: alice,
		Index:   "1",
		Code:    "test",
	})
	require.Nil(t, err)
	require.EqualValues(t, dataproc.MsgSetCodeResponse{
		CodeIndex: "0",
	}, *createResponse)

	createResponse, err = msgServer.SetCode(context, &dataproc.MsgSetCode{
		Creator: alice,
		Index:   "1",
		Code:    "test",
	})
	require.Nil(t, err)
	require.EqualValues(t, dataproc.MsgSetCodeResponse{
		CodeIndex: "1",
	}, *createResponse)
}
