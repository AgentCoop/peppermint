package join

import (
	"context"
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinHelloRequest struct {
	server.RequestHeader
	nodePubKey []byte
}

type DataBag interface {
	NodePubKey() []byte
}

func NewJoinHello(ctx context.Context, original *msg.JoinHello_Request) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.RequestHeader = server.NewRequestHeader(ctx)
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
	server.Response
	hubPubKey []byte
}

func NewJoinHelloResponse(hubPubKey []byte) *joinHelloResponse {
	r := new(joinHelloResponse)
	r.Response = server.NewResponseHeader()
	r.hubPubKey = hubPubKey
	return r
}

func (r *joinHelloResponse) ToGrpcResponse() interface{} {
	resp := new(msg.JoinHello_Response)
	resp.DhPubKey = r.hubPubKey
	return resp
}
