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

func NewJoinHello(orig *msg.JoinHello_Request) *joinHelloRequest {
	r := new(joinHelloRequest)
	r.nodePubKey = orig.GetDhPubKey()
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

type joinHelloResponse struct {
	resp *msg.JoinHello_Response
}

func NewJoinHelloResponse(pubKey []byte) joinHelloResponse {
	r := joinHelloResponse{}
	r.resp = &msg.JoinHello_Response{DhPubKey: pubKey}
	return r
}

func (r joinHelloResponse) ToGrpc() interface{} {
	return r.resp
}
