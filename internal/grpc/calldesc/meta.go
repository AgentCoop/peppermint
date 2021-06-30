package calldesc

import (
	i "github.com/AgentCoop/peppermint/internal"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func (m *meta) copySessionId(src *clDescriptor) {
	m.sId = src.meta.SessionId()
	m.SetSessionId(src.meta.SessionId())
}

func (m *meta) extractCommonFieldsVals() {
	m.sId = g.ExtractSessionId(&m.header)
	m.nodeId = g.ExtractNodeId(&m.header)
}

func (m *meta) SessionId() i.SessionId {
	return m.sId
}

func (m *meta) SetSessionId(id i.SessionId) {
	g.SetSessionId(&m.header, id)
}

func (m *meta) NodeId() i.NodeId {
	return m.nodeId
}

func (m *meta) SetHeader(newMd metadata.MD) {
	merged := metadata.Join(m.header, newMd)
	m.header = merged
}

func (m *meta) SendHeader(metadata.MD) error {
	switch m.parent.typ {
	case ServerType:
		err := grpc.SendHeader(m.parent.Context, m.header)
		return err
	case ClientType:
		m.parent.Context = metadata.NewOutgoingContext(m.parent.Context, m.header)
		err := grpc.SendHeader(m.parent.Context, m.header)
		return err
	}
	return nil
}

func (m *meta) Header() *metadata.MD {
	return &m.header
}

func (m *meta) Trailer() *metadata.MD {
	return &m.trailer
}

func (m *meta) SetTrailer(md metadata.MD) {
	panic("SetTrailer implement me")
}
