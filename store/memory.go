package store

import "fmt"

type memory struct {
	storage map[string]string // ID -> Value
	revstor map[string]string // Value -> ID
}

func makeMemoryStore() *memory {
	s := make(map[string]string)
	r := make(map[string]string)
	return &memory{storage: s, revstor: r}
}

// getID returns id if the value is in store.
// Otherwise it returns an empty string.
func (m memory) getID(value string) string {
	return m.revstor[value]
}

func (m memory) get(id string) string {
	return m.storage[id]
}

func (m memory) put(id, value string) error {
	if id == "" {
		return fmt.Errorf("Empty id is not allowed")
	}
	m.storage[id] = value
	m.revstor[value] = id
	return nil
}
