package keeper_test

// import (
// 	"context"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/require"

// 	keepertest "github.com/igor-sikachyna/dataproc/testutil/keeper"
// 	"github.com/igor-sikachyna/dataproc/x/dataproc/keeper"
// 	dataproc "github.com/igor-sikachyna/dataproc/x/dataproc/module"
// 	"github.com/igor-sikachyna/dataproc/x/dataproc/testutil"
// 	"github.com/igor-sikachyna/dataproc/x/dataproc/types"
// )

// func setupMsgServerSetCode(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
// 	server, k, context, _, escrow := setupMsgServerCreateGameWithMock(t)
// 	escrow.ExpectAny(context)
// 	return server, k, context
// }

// func setupMsgServerSetCodeWithMock(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context,
// 	*gomock.Controller, *testutil.MockBankKeeper) {
// 	ctrl := gomock.NewController(t)
// 	bankMock := testutil.NewMockBankKeeper(ctrl)
// 	k, ctx := keepertest.CheckersKeeperWithMocks(t, bankMock)
// 	dataproc.InitGenesis(ctx, k, *types.DefaultGenesis())
// 	server := keeper.NewMsgServerImpl(k)
// 	return server, k, ctx, ctrl, bankMock
// }
