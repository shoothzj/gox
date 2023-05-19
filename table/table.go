package table

// Table is a generic two-level map
type Table[K1 comparable, K2 comparable, V any] map[K1]map[K2]V
