package join

import msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"

type joinRequest struct {
	secret string
	tags []string
}

func NewJoin(original *msg.Join_Request) *joinRequest {
	r := new(joinRequest)
	r.secret = original.GetJoinSecret()
	r.tags = original.GetTag()
	return r
}

func (v *joinRequest) Run() error {
	return nil
}

func (r *joinRequest) Validate() error {
	return nil
}
