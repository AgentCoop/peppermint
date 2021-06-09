package runtime

type Store_FallbackHandler func(key interface{}) (interface{}, error)

type Store interface {
	RegisterFallback(handler Store_FallbackHandler)
}

type Backend interface {
	Has(key interface{}) bool
	Get(key interface{}) (interface{}, error)
	Set(key interface{}, val interface{})
	Clear()
}

type InMemoryStore interface {
	Store
	Backend
}
