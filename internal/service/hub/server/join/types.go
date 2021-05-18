package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinCtx struct {
	encKey []byte
	ReqChan [1]server.PairChan
	ResChan [1]server.PairChan
}

func NewJoinContext() *joinCtx {
	ctx := new(joinCtx)
	ctx.ReqChan = [1]server.PairChan{
		make(server.PairChan),
	}
	ctx.ResChan = [1]server.PairChan{
		make(server.PairChan),
	}
	//ctx.JoinHelloReqCh = make(server.PairChan)
	//ctx.JoinHelloRespCh = make(server.PairChan)
	return ctx
}