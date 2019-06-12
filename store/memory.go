package store

import (
	"encoding/base64"
	"strconv"
	"sync"
)

var storeIndex int

type memory struct {
	storage  map[string]string // ID -> Value
	revstor  map[string]string // Value -> ID
	encoding *base64.Encoding
	mutex    sync.RWMutex
}

func makeMemoryStore() *memory {
	s := make(map[string]string)
	r := make(map[string]string)
	e := base64.URLEncoding.WithPadding(base64.NoPadding)
	return &memory{storage: s, revstor: r, encoding: e}
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

func (m *memory) put(value string) (string, error) {
	m.mutex.Lock()
	storeIndex++
	id := m.encoding.EncodeToString([]byte(strconv.Itoa(storeIndex)))
	m.storage[id] = value
	m.revstor[value] = id
	m.mutex.Unlock()
	return id, nil
}
