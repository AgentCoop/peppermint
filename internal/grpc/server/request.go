package server

import (
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/metadata"
	"context"
)

type requestHeader struct {
	sessId grpc.SessionId
	nodeId grpc.NodeId
}

type RequestHeader interface {
	SessionId() grpc.SessionId
	NodeId() grpc.NodeId
}

type RequestData interface {
	Populate(original interface{})
	Validate() error
}

type Request interface {
	RequestHeader
	RequestData
}

func NewRequestHeader(ctx context.Context) *requestHeader {
	r := new(requestHeader)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return r
	}

	var vals []string
	vals = md.Get(grpc.META_FIELD_NODE_ID)
	if len(vals) > 1 {
		r.nodeId = grpc.NodeId(utils.Hex2int(vals[0]))
	}
	return r
}

func (r *requestHeader) SessionId() grpc.SessionId {
	return r.sessId
}

func (r *requestHeader) NodeId() grpc.NodeId {
	return r.nodeId
}
