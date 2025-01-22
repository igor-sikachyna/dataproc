package dataproc

import "cosmossdk.io/collections"

const ModuleName = "dataproc"
const MaxIndexLength = 256

var (
	ParamsKey      = collections.NewPrefix("Params")
	SystemInfoKey  = collections.NewPrefix("SystemInfo")
	StoredCodesKey = collections.NewPrefix("StoredCodes/value/")
)
