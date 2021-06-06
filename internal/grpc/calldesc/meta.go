package calldesc

import (
	i "github.com/AgentCoop/peppermint/internal"
	"google.golang.org/grpc/metadata"
)

func (m *meta) SessionId() i.SessionId {
	return m.sId
}

func (m *meta) NodeId() i.NodeId {
	return m.nodeId
}

func (m *meta) SetHeader(metadata.MD) {

}

func (m *meta) SendHeader(metadata.MD) error {
	return nil
}

func (m *meta) Header() *metadata.MD {
	return &m.header
}

func (m *meta) Trailer() *metadata.MD {
	return &m.trailer
}
