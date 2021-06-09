package store

func NewInMemoryStore() *inMemStore {
	s := new(inMemStore)
	s.backend = make(inMemBackend, 0)
	return s
}
