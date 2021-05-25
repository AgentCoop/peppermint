package join

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/client"
)

type joinHelloRequest struct {
	client.Request
	pubKey []byte
}

func (r *joinHelloRequest) SetSessionId(id g.SessionId) {
	panic("implement me")
}

//func (r *joinHelloRequest) SendHeader() {
//	panic("implement me")
//}

func NewJoinHello(pair client.RequestResponsePair, pubKey []byte) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.pubKey = pubKey
	r.Request = pair.AssignNewRequest(r)
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
	client.Response
	original *hub.JoinHello_Response
	hubPubKey []byte
}

func NewJoinHelloResponse(pair client.RequestResponsePair, original *hub.JoinHello_Response) *joinHelloResponse {
	r := new(joinHelloResponse)
	r.original = original
	r.Populate()
	r.Response = pair.AssignNewResponse(r)
	return r
}

func (r *joinHelloResponse) Populate() {
	r.hubPubKey = r.original.GetDhPubKey()
}

func (r *joinHelloResponse) HubPubKey() []byte {
	return r.hubPubKey
}
