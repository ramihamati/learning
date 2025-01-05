package types

const (
	// ModuleName defines the module name
	ModuleName = "legal"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_legal"
)

var (
	ParamsKey = []byte("p_legal")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
