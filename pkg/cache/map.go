package cache

import (
	"sync"
	"time"
)

// Map ...
type Map struct {
	m           map[string]time.Time
	mutex       sync.RWMutex
	expireAfter time.Duration
}

// NewMap creates a new cache map
func NewMap(expireAfter time.Duration) *Map {
	return &Map{
		m:           map[string]time.Time{},
		expireAfter: expireAfter,
	}
}

// Store saves key in a map
// returns true if key not in cache or expired
func (m *Map) Store(k string) bool {
	// attempt to retreive the record
	m.mutex.RLock()
	expires, ok := m.m[k]
	m.mutex.RUnlock()
	if ok {
		// check if timeframe has expired
		if expires.After(time.Now()) {
			return false
		}
	}

	// assuming record cannot be retreived, store it
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.m[k] = time.Now().Add(m.expireAfter)
	return true
}
