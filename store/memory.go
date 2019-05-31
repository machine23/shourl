package store

type memory struct {
	storage map[string]string
}

func makeMemoryStore() *memory {
	s := make(map[string]string)
	return &memory{storage: s}
}

func (m memory) get(id string) string {
	return m.storage[id]
}

func (m memory) put(id, value string) error {
	m.storage[id] = value
	return nil
}
