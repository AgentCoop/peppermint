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
	joinReqCh client.ReqChan
	joinResCh client.ResChan
}

func NewJoinContext(address string, secret string) *joinCtx {
	ctx := new(joinCtx)
	ctx.HubClient = hub.NewClient(client.NewBaseClient(address))
	ctx.secret = secret
	ctx.joinHelloReqCh = make(client.ReqChan)
	ctx.joinHelloResCh = make(client.ResChan)
	ctx.joinReqCh = make(client.ReqChan)
	ctx.joinResCh = make(client.ResChan)
	return ctx
}

func (j *joinCtx) JoinHelloRequest() client.ReqChan {
	return j.joinHelloReqCh
}

func (j *joinCtx) JoinHelloResponse() client.ResChan {
	return j.joinHelloResCh
}

func (j *joinCtx) JoinRequest() client.ReqChan {
	return j.joinReqCh
}

func (j *joinCtx) JoinResponse() client.ResChan {
	return j.joinResCh
}
