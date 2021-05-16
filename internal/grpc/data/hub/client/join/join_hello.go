package join

import "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"

type joinHelloRequest struct {
	pubKey []byte
}

func NewJoinHello(pubKey []byte) *joinHelloRequest {
	r := new(joinHelloRequest)
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
	original *hub.JoinHello_Response
	hubPubKey []byte
}

func NewJoinHelloResponse(original *hub.JoinHello_Response) *joinHelloResponse {
	r := new(joinHelloResponse)
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
