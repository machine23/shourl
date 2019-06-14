package store

import (
	"fmt"
	"sync"
)

type memory struct {
	storage map[string]string // ID -> Value
	revstor map[string]string // Value -> ID
	mutex   sync.RWMutex
}

func makeMemoryStore() *memory {
	s := make(map[string]string)
	r := make(map[string]string)
	return &memory{storage: s, revstor: r}
}

// getID returns id if the value is in store.
// Otherwise it returns an empty string.
func (m *memory) getID(value string) string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.revstor[value]
}

func (m *memory) get(id string) string {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	return m.storage[id]
}

func (m *memory) put(id, value string) error {
	if id == "" {
		return fmt.Errorf("Empty id is not allowed")
	}
	m.mutex.Lock()
	m.storage[id] = value
	m.revstor[value] = id
	m.mutex.Unlock()
	return nil
}
