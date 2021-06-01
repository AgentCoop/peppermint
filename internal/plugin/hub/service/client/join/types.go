package join

import (
	"github.com/AgentCoop/peppermint/internal/grpc/client"
	client2 "github.com/AgentCoop/peppermint/internal/service/hub/grpc/client"
	"net"
)

type joinCtx struct {
	client2.HubClient
	encKey []byte
	secret string
	reqChan []client.ReqChan
	resChan []client.ResChan
}

func NewJoinContext(address net.Addr, secret string) *joinCtx {
	ctx := new(joinCtx)
	ctx.HubClient = client2.NewClient(client.NewBaseClient(address))
	ctx.secret = secret
	// Init channels
	ctx.reqChan = []client.ReqChan{
		make(client.ReqChan, 0),
		make(client.ReqChan, 0),
	}
	ctx.resChan = []client.ResChan{
		make(client.ResChan, 0),
		make(client.ResChan, 0),
	}
	return ctx
}

func (j *joinCtx) ReqChan(idx int) client.ReqChan {
	return j.reqChan[idx]
}

func (j *joinCtx) ResChan(idx int) client.ResChan {
	return j.resChan[idx]
}
