package join

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
)

type joinRequest struct {
	server.Request
	secret string
	tags []string
}

type Join_DataBag interface {
	Secret() string
	Tags() []string
}

func NewJoin(pair server.RequestResponsePair, original *msg.Join_Request) *joinRequest {
	r := new(joinRequest)
	r.secret = original.JoinSecret
	r.tags  = original.Tag
	r.Request = pair.AssignNewRequest(r)
	return r
}

func (r *joinRequest) Validate() error {
	return nil
}

func (r *joinRequest) Secret() string {
	return r.secret
}

func (r *joinRequest) Tags() []string {
	return r.tags
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
