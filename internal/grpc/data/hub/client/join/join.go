package join

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
)

type joinRequest struct {
	client.Request
	secret string
}

//func (r *joinHelloRequest) SetSessionId(id g.SessionId) {
//	panic("implement me")
//}
//
//func (r *joinHelloRequest) SendHeader() {
//	panic("implement me")
//}

func NewJoin(ctx context.Context, secret string) *joinRequest {
	r := new(joinRequest)
	r.Request = client.NewRequest(ctx)
	r.secret = secret
	return r
}

func (r *joinRequest) ToGrpcRequest() interface{} {
	greq :=  new(hub.Join_Request)
	return greq
}

//
// Responses
//

type Join_DataBag interface {
	Secret() []byte
}

type joinResponse struct {
	client.Response
	original *hub.Join_Response
}

func NewJoinResponse(ctx context.Context, original *hub.Join_Response) *joinResponse {
	r := new(joinResponse)
	r.original = original
	r.Response = client.NewResponse(ctx)
	r.Populate()
	return r
}

func (r *joinResponse) Populate() {
}

