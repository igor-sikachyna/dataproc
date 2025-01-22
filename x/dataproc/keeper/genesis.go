package keeper

import (
	"context"

	"github.com/igor-sikachyna/dataproc/x/dataproc"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *dataproc.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, indexedStoredCode := range data.IndexedStoredCodeList {
		if err := k.StoredCodes.Set(ctx, indexedStoredCode.Index, indexedStoredCode.StoredCode); err != nil {
			return err
		}
	}

	if err := k.SetSystemInfo(ctx, data.SystemInfo); err != nil {
		return err
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*dataproc.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var indexedStoredCodes []dataproc.IndexedStoredCode
	if err := k.StoredCodes.Walk(ctx, nil, func(index string, storedCode dataproc.StoredCode) (bool, error) {
		indexedStoredCodes = append(indexedStoredCodes, dataproc.IndexedStoredCode{
			Index:      index,
			StoredCode: storedCode,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &dataproc.GenesisState{
		Params:                params,
		IndexedStoredCodeList: indexedStoredCodes,
	}, nil
}
