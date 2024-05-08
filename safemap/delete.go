package safemap

// Deleting key value
func (sm *SafeMap[T, V]) DeleteFromSafeMap(key T, value V) {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Deletion of key value entry from safemap
	delete(sm.data, key)
}
