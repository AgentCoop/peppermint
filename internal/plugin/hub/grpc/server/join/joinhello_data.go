package join

import (
	msg "github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
)

type joinHelloRequest struct {
	nodePubKey []byte
}

type joinHello_DataBag interface {
	NodePubKey() []byte
}

func NewJoinHello(original *msg.JoinHello_Request) *joinHelloRequest {
	r := new(joinHelloRequest)
	return r
}

func (v *joinHelloRequest) Run() error {
	return nil
}

func (r *joinHelloRequest) Validate() error {
	return nil
}

func (r *joinHelloRequest) NodePubKey() []byte {
	return r.nodePubKey
}
