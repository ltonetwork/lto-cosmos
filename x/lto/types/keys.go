package types

const (
	// ModuleName defines the module name
	ModuleName = "lto"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lto"

	// keep track of anchor indexes
	AnchorKey      = "Anchor/value/"
	AnchorCountKey = "Anchor/count/"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
