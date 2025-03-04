package keeper

import (
	"context"
	"errors"

	"cosmossdk.io/collections"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/igor-sikachyna/dataproc/x/dataproc"
)

var _ dataproc.QueryServer = queryServer{}

// NewQueryServerImpl returns an implementation of the module QueryServer.
func NewQueryServerImpl(k Keeper) dataproc.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}

// GetCode defines the handler for the Query/GetCode RPC method.
func (qs queryServer) GetCode(ctx context.Context, req *dataproc.QueryGetCodeRequest) (*dataproc.QueryGetCodeResponse, error) {
	code, err := qs.k.StoredCodes.Get(ctx, req.Index)
	if err == nil {
		return &dataproc.QueryGetCodeResponse{Code: &code}, nil
	}
	if errors.Is(err, collections.ErrNotFound) {
		return &dataproc.QueryGetCodeResponse{Code: nil}, nil
	}

	return nil, status.Error(codes.Internal, err.Error())
}
