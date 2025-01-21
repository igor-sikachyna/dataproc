package types

const (
	// ModuleName defines the module name
	ModuleName = "dataproc"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	//
)

var (
	ParamsKey = []byte("p_dataproc")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	Bech32CodecPrefix = "cosmos"
)
