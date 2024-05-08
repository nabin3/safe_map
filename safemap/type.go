package safemap

import "sync"

// BluePrint of safemap
type SafeMap[T comparable, V any] struct {
	data map[T]V
	mu   sync.RWMutex
}
