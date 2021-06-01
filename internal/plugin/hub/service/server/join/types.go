package join

type joinCtx struct {
	encKey []byte
}

func NewJoinContext() *joinCtx {
	ctx := new(joinCtx)
	return ctx
}
