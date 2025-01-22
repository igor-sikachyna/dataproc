package keeper

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"cosmossdk.io/collections"
	"github.com/igor-sikachyna/dataproc/x/dataproc"
)

type msgServer struct {
	k Keeper
}

var _ dataproc.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the module MsgServer interface.
func NewMsgServerImpl(keeper Keeper) dataproc.MsgServer {
	return &msgServer{k: keeper}
}

// SetCode defines the handler for the MsgSetCode message.
func (ms msgServer) SetCode(ctx context.Context, msg *dataproc.MsgSetCode) (*dataproc.MsgSetCodeResponse, error) {
	if length := len([]byte(msg.Index)); dataproc.MaxIndexLength < length || length < 1 {
		return nil, dataproc.ErrIndexTooLong
	}
	if _, err := ms.k.StoredCodes.Get(ctx, msg.Index); err == nil || errors.Is(err, collections.ErrEncoding) {
		return nil, fmt.Errorf("code already exists at index: %s", msg.Index)
	}

	systemInfo, err := ms.k.GetSystemInfo(ctx)
	if err != nil {
		return nil, err
	}
	newIndex := strconv.FormatUint(systemInfo.NextId, 10)

	// newBoard := rules.New()
	// storedGame := dataproc.StoredGame{
	// 	Board: newBoard.String(),
	// 	Turn:  rules.PieceStrings[newBoard.Turn],
	// 	Black: msg.Black,
	// 	Red:   msg.Red,
	// }
	// if err := storedGame.Validate(); err != nil {
	// 	return nil, err
	// }
	// if err := ms.k.StoredGames.Set(ctx, msg.Index, storedGame); err != nil {
	// 	return nil, err
	// }

	ms.k.SetSystemInfo(ctx, dataproc.SystemInfo{
		NextId: systemInfo.NextId + 1,
	})

	return &dataproc.MsgSetCodeResponse{
		CodeIndex: newIndex,
	}, nil
}
