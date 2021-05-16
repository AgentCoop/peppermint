package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	"github.com/AgentCoop/peppermint/internal/grpc/client/hub"
)

type joinCtx struct {
	hub.HubClient
	encKey []byte
	secret string
	joinHelloReqCh client.ReqChan
	joinHelloResCh client.ResChan
	joinReqCh chan client.Request
	joinResCh chan client.Response
}

func NewJoinContext(address string, secret string) *joinCtx {
	ctx := new(joinCtx)
	ctx.HubClient = hub.NewClient(client.NewBaseClient(address))
	ctx.secret = secret
	ctx.joinHelloReqCh = make(client.ReqChan)
	ctx.joinHelloResCh = make(client.ResChan)
	return ctx
}

func (j *joinCtx) JoinHelloRequest() client.ReqChan {
	return j.joinHelloReqCh
}

func (j *joinCtx) JoinHelloResponse() client.ResChan {
	return j.joinHelloResCh
}
