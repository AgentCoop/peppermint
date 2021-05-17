package join

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinRequest struct {
	server.MetaData
	nodePubKey []byte
}

type Join_DataBag interface {
}

func NewHello(md server.MetaData, original *msg.Join_Request) *joinRequest {
	r := new(joinRequest)
	r.MetaData = md
	r.Populate(original)
	return r
}

func (r *joinRequest) Populate(original interface{}) {
}

func (r *joinRequest) Validate() error {
	return nil
}

//
// Response
//

type joinResponse struct {
	server.MetaData
}

func NewJoinResponse(md server.MetaData, hubPubKey []byte) *joinResponse {
	r := new(joinResponse)
	r.MetaData = md
	return r
}

func (r *joinResponse) ToGrpcResponse() interface{} {
	resp := new(msg.Join_Response)
	return resp
}
