package chello

type clientHelloCtx struct {
	encKey  []byte
	randMsg []byte
}

func NewContext() *clientHelloCtx {
	return new(clientHelloCtx)
}

func (c *clientHelloCtx) EncKey() []byte {
	return c.encKey
}
