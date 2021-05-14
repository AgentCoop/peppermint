package join

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinHelloRequest struct {
	server.RequestHeader
	nodePubKey []byte
}

type DataBag interface {
	NodePubKey() []byte
}

func NewJoinHello(ctx context.Context, original *hub.JoinHello_Request) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.RequestHeader = server.NewRequestHeader(ctx)
	r.Populate(original)
	return r
}

func (r *joinHelloRequest) Populate(original interface{}) {
	r.nodePubKey = original.(*hub.JoinHello_Request).GetDhPubKey()
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

func NewJoinHelloResponse(ctx context.Context) *joinHelloResponse {
	r := new(joinHelloResponse)
	r.Response = server.NewResponseHeader(ctx)
	return r
}

func (r *joinHelloResponse) ToGrpcResponse() interface{} {
	panic("implement me")
}
