package types

const (
	// ModuleName defines the module name
	ModuleName = "distribute"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_distribute"
)

var (
	ParamsKey = []byte("p_distribute")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
