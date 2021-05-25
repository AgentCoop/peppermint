package join

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
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

func NewJoin(pair client.RequestResponsePair, secret string) *joinRequest {
	r := new(joinRequest)
	r.secret = secret
	r.Request = pair.AssignNewRequest(r)
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

func NewJoinResponse(pair client.RequestResponsePair, original *hub.Join_Response) *joinResponse {
	r := new(joinResponse)
	r.original = original
	r.Response = pair.AssignNewResponse(r)
	r.Populate()
	return r
}

func (r *joinResponse) Populate() {
}

