package join

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
)

type joinHelloRequest struct {
	client.Request
	pubKey []byte
}

func (r *joinHelloRequest) SetSessionId(id g.SessionId) {
	//panic("implement me")
}

func (r *joinHelloRequest) SendHeader() {
	//panic("implement me")
}

func NewJoinHello(ctx context.Context, pubKey []byte) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.Request = client.NewRequest(ctx)
	r.pubKey = pubKey
	return r
}

func (r *joinHelloRequest) ToGrpcRequest() interface{} {
	greq :=  new(hub.JoinHello_Request)
	greq.DhPubKey = r.pubKey
	return greq
}

//
// Responses
//

type JoinHello_DataBag interface {
	HubPubKey() []byte
}

type joinHelloResponse struct {
	context.Context
	client.ResponseHeader
	original *hub.JoinHello_Response
	hubPubKey []byte
}

func NewJoinHelloResponse(ctx context.Context, original *hub.JoinHello_Response) *joinHelloResponse {
	r := new(joinHelloResponse)
	r.Context = ctx
	r.ResponseHeader = client.NewResponse(ctx)
	r.original = original
	r.Populate()
	return r
}

func (r *joinHelloResponse) Populate() {
	r.hubPubKey = r.original.GetDhPubKey()
}

func (r *joinHelloResponse) HubPubKey() []byte {
	return r.hubPubKey
}
