package dataproc

// NewGenesisState creates a new genesis state with default values.
func NewGenesisState() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
func (gs *GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	unique := make(map[string]bool)
	for _, indexedStoredCode := range gs.IndexedStoredCodeList {
		if length := len([]byte(indexedStoredCode.Index)); MaxIndexLength < length || length < 1 {
			return ErrIndexTooLong
		}
		if _, ok := unique[indexedStoredCode.Index]; ok {
			return ErrDuplicateAddress
		}
		if err := indexedStoredCode.StoredCode.Validate(); err != nil {
			return err
		}
		unique[indexedStoredCode.Index] = true
	}

	return nil
}
