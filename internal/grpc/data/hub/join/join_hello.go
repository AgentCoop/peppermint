package join

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinHelloRequest struct {
	server.Request
	nodePubKey []byte
}

type DataBag interface {
	NodePubKey() []byte
}

func NewJoinHello(md server.MetaData, original *msg.JoinHello_Request) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.Request = md
	r.Populate(original)
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
	server.MetaData
	hubPubKey []byte
}

func NewJoinHelloResponse(md server.MetaData, hubPubKey []byte) *joinHelloResponse {
	r := new(joinHelloResponse)
	r.MetaData = md
	r.hubPubKey = hubPubKey
	return r
}

func (r *joinHelloResponse) ToGrpcResponse() interface{} {
	resp := new(msg.JoinHello_Response)
	resp.DhPubKey = r.hubPubKey
	return resp
}
