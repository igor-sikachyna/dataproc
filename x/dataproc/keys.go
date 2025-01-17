package dataproc

import "cosmossdk.io/collections"

const ModuleName = "dataproc"
const MaxIndexLength = 256

var (
	ParamsKey      = collections.NewPrefix("Params")
	StoredGamesKey = collections.NewPrefix("StoredGames/value/")
)
