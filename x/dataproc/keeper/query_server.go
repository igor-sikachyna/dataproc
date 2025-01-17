package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/igor-sikachyna/dataproc"
)

var _ dataproc.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) dataproc.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// GetGame defines the handler for the Query/GetGame RPC method.
func (qs queryServer) GetGame(ctx context.Context, req *dataproc.QueryGetGameRequest) (*dataproc.QueryGetGameResponse, error) {
	game, err := qs.k.StoredGames.Get(ctx, req.Index)
	if err == nil {
		return &dataproc.QueryGetGameResponse{Game: &game}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &dataproc.QueryGetGameResponse{Game: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
