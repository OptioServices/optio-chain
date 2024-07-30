package types

const (
	// ModuleName defines the module name
	ModuleName = "rewardtoken"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_rewardtoken"
)

var (
	ParamsKey = []byte("p_rewardtoken")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
