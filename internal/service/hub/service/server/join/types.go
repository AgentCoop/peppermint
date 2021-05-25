package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinCtx struct {
	encKey []byte
	reqChan [2]server.PairChan
	resChan [2]server.PairChan
}

func NewJoinContext() *joinCtx {
	ctx := new(joinCtx)
	ctx.reqChan = [2]server.PairChan{
		make(server.PairChan),
		make(server.PairChan),
	}
	ctx.resChan = [2]server.PairChan{
		make(server.PairChan),
		make(server.PairChan),
	}
	return ctx
}

func (ctx *joinCtx) ReqChan() [2]server.PairChan {
	return ctx.reqChan
}

func (ctx *joinCtx) ResChan() [2]server.PairChan {
	return ctx.resChan
}