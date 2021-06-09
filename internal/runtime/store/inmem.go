package store

type inMemBackend map[interface{}]interface{}

type inMemStore struct {
	store
	backend inMemBackend
}

func (s *inMemStore) Has(key interface{}) bool {
	_, ok := s.backend[key]
	return ok
}

func (s *inMemStore) Get(key interface{}) (interface{}, error) {
	var (
		val interface{}
		err error
		ok bool
	)
	if val, ok =  s.backend[key]; !ok {
		if val, err = s.fallback(key); err != nil {
			return nil, err
		}
	}
	s.backend[key] = val
	return val, nil
}

func (s *inMemStore) Set(key interface{}, val interface{}) {
	s.backend[key] = val
}

func (s *inMemStore) Clear() {
	s.backend = make(map[interface{}]interface{}, 0)
}
