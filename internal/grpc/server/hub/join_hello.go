package hub

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"

	"context"
)

func (s *server) JoinHello(c context.Context, req *hub.JoinHello_Request) (*hub.JoinHello_Response, error) {
	res := &hub.JoinHello_Response{
		CryptoNonce: nil,
		DfPubKey:    nil,
	}
	return res, nil
}