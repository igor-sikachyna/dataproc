package keeper

import (
	"context"

	"github.com/igor-sikachyna/dataproc/x/dataproc"
)

// GetSystemInfo returns the module's system information.
func (k Keeper) GetSystemInfo(ctx context.Context) (dataproc.SystemInfo, error) {
	return k.SystemInfo.Get(ctx)
}

// GetSystemInfo returns the module's system information.
func (k Keeper) HasSystemInfo(ctx context.Context) (dataproc.SystemInfo, error) {
	return k.SystemInfo.Get(ctx)
}

// SetSystemInfo sets the module's system information.
func (k *Keeper) SetSystemInfo(ctx context.Context, systemInfo dataproc.SystemInfo) {
	k.SystemInfo.Set(ctx, systemInfo)
}
