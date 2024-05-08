package safemap

import "sync"

// Creating a new safemap
func NewSafeMap[T comparable, V any]() *SafeMap[T, V] {
	return &SafeMap[T, V]{
		data: make(map[T]V),
		mu:   sync.RWMutex{},
	}
}
