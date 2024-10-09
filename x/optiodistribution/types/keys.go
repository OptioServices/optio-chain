package types

const (
	// ModuleName defines the module name
	ModuleName = "optiodistribution"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_optiodistribution"
)

var (
	ParamsKey = []byte("p_optiodistribution")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
