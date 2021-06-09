package runtime

type FallbackHandler func(key interface{}) (interface{}, error)

type Store interface {
	Has(key interface{}) bool
	Get(key interface{}) (interface{}, error)
	GetFallback(handler FallbackHandler)
	Set(key interface{}, val interface{})
	Clear()
}
