package safemap

import "errors"

// Insertion of new Key Value
func (sm *SafeMap[T, V]) AddNewKeyValueToSafeMap(key T, value V) error {
	sm.mu.Lock()
	defer sm.mu.Unlock()

	// Checking if given key already exists
	if _, exist := sm.data[key]; exist {
		return errors.New("given key already existes, didn't add the given key value")
	}

	// Adding new key value on map
	sm.data[key] = value
	return nil
}
