package chello

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/backoffice/hub"
)

type clientHelloRequest struct {
	nodePubKey []byte
}

type clientHelloData interface {
	NodePubKey() []byte
}

func NewClientHelloRequest(orig *hub.ClientHello_Request) *clientHelloRequest {
	r := new(clientHelloRequest)
	r.nodePubKey = orig.GetDhPubKey()
	return r
}

func (v *clientHelloRequest) Run() error {
	return nil
}

func (r *clientHelloRequest) Validate() error {
	return nil
}

func (r *clientHelloRequest) NodePubKey() []byte {
	return r.nodePubKey
}

type clientHelloResponse struct {
	resp *hub.ClientHello_Response
}

func NewClientHelloResponse(pubKey []byte, randMsg []byte) *clientHelloResponse {
	r := new(clientHelloResponse)
	r.resp = &hub.ClientHello_Response{
		DhPubKey: pubKey,
		RandMsg: randMsg,
	}
	return r
}

func (r clientHelloResponse) ToGrpc() interface{} {
	return r.resp
}
