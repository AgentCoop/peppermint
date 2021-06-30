package chello

import msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"

type joinRequest struct {
	secret string
	Tags   []string
}

func NewJoin(original *msg.Join_Request) *joinRequest {
	r := new(joinRequest)
	r.secret = original.GetJoinSecret()
	r.Tags = original.GetTag()
	return r
}

func (v *joinRequest) Run() error {
	return nil
}

func (r *joinRequest) Validate() error {
	return nil
}

type joinResponse struct {
	resp *msg.Join_Response
}

func NewJoinResponse() *joinResponse {
	r := new(joinResponse)
	r.resp = &msg.Join_Response{}
	return r
}

func (r joinResponse) ToGrpc() interface{} {
	return r.resp
}

