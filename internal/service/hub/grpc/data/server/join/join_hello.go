package join

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinHelloRequest struct {
	server.Request
	nodePubKey []byte
}

type DataBag interface {
	NodePubKey() []byte
}

func NewJoinHello(pair server.RequestResponsePair, original *msg.JoinHello_Request) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.Populate(original)
	r.Request = pair.AssignNewRequest(r)
	return r
}

func (r *joinHelloRequest) Populate(original interface{}) {
	r.nodePubKey = original.(*msg.JoinHello_Request).GetDhPubKey()
}

func (r *joinHelloRequest) Validate() error {
	return nil
}

func (r *joinHelloRequest) NodePubKey() []byte {
	return r.nodePubKey
}

//
// Response
//

type joinHelloResponse struct {
	server.Response
	hubPubKey []byte
}

func NewJoinHelloResponse(pair server.RequestResponsePair, hubPubKey []byte) *joinHelloResponse {
	r := new(joinHelloResponse)
	r.hubPubKey = hubPubKey
	r.Response = pair.AssignNewResponse(r)
	return r
}

func (r *joinHelloResponse) ToGrpcResponse() interface{} {
	resp := new(msg.JoinHello_Response)
	resp.DhPubKey = r.hubPubKey
	return resp
}
