package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinCtx struct {
	encKey []byte
	JoinHelloReqCh chan server.Request
	JoinHelloRespCh chan server.Response
}

func NewJoinContext() *joinCtx {
	ctx := new(joinCtx)
	ctx.JoinHelloReqCh = make(chan server.Request)
	ctx.JoinHelloRespCh = make(chan server.Response)
	return ctx
}