package hub

import (
	"context"
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service/hub"
	//"github.com/AgentCoop/peppermint/internal/api/peppermint/service"
)

func (c *hubClient) JoinHello(dfPublicKey []byte) {
	r := &hub.JoinHello_Request{DfPubKey: dfPublicKey}
	c.grpcHandle.JoinHello(context.Background(), r)
}
