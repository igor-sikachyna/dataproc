package keeper

import (
	"context"

	"github.com/igor-sikachyna/dataproc"
)

// InitGenesis initializes the module state from a genesis state.
func (k *Keeper) InitGenesis(ctx context.Context, data *dataproc.GenesisState) error {
	if err := k.Params.Set(ctx, data.Params); err != nil {
		return err
	}

	for _, indexedStoredGame := range data.IndexedStoredGameList {
		if err := k.StoredGames.Set(ctx, indexedStoredGame.Index, indexedStoredGame.StoredGame); err != nil {
			return err
		}
	}

	return nil
}

// ExportGenesis exports the module state to a genesis state.
func (k *Keeper) ExportGenesis(ctx context.Context) (*dataproc.GenesisState, error) {
	params, err := k.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	var indexedStoredGames []dataproc.IndexedStoredGame
	if err := k.StoredGames.Walk(ctx, nil, func(index string, storedGame dataproc.StoredGame) (bool, error) {
		indexedStoredGames = append(indexedStoredGames, dataproc.IndexedStoredGame{
			Index:      index,
			StoredGame: storedGame,
		})
		return false, nil
	}); err != nil {
		return nil, err
	}

	return &dataproc.GenesisState{
		Params:                params,
		IndexedStoredGameList: indexedStoredGames,
	}, nil
}
