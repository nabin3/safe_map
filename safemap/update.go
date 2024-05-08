package safemap

// Updating value of given key
func (sm *SafeMap[T, V]) UpdateSafeMap(key T, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Updating existing key's value or creating new key value
	sm.data[key] = value
}
