package safemap

import "errors"

// Insertion of new Key Value
func (sm *SafeMap[T, V]) GetValueFromSafeMap(key T) (any, error) {
	sm.mu.RLock()
	defer sm.mu.RUnlock()

	var defaultValue any
	// Checking if given key exists or not
	if value, exist := sm.data[key]; !exist {
		return defaultValue, errors.New("given key doesn't exist")
	} else {
		return value, nil
	}

}
