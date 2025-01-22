package keeper

import (
	"context"

	"github.com/igor-sikachyna/dataproc/x/dataproc"
)

// GetSystemInfo returns the module's system information.
func (k Keeper) GetSystemInfo(ctx context.Context) (dataproc.SystemInfo, error) {
	return k.SystemInfo.Get(ctx)
}

// HasSystemInfo returns whether module's system information exists.
func (k Keeper) HasSystemInfo(ctx context.Context) (bool, error) {
	return k.SystemInfo.Has(ctx)
}

// SetSystemInfo sets the module's system information.
func (k *Keeper) SetSystemInfo(ctx context.Context, systemInfo dataproc.SystemInfo) error {
	return k.SystemInfo.Set(ctx, systemInfo)
}
