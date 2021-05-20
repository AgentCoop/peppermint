package server

import (
	"github.com/AgentCoop/peppermint/internal/utils"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/metadata"
	"context"
)

type request struct {
	context.Context
	sessId grpc.SessionId
	nodeId grpc.NodeId
}

type RequestHeader interface {
	SessionId() grpc.SessionId
	NodeId() grpc.NodeId
	RequiredServices() []string // Optional. If presented, hub must forward request to a host with the required services available
}

type RequestData interface {
	Populate(original interface{})
	Validate() error
}

type Request interface {
	//context.Context
	RequestHeader
	//RequestData
}

func NewRequest(ctx context.Context) *request {
	r := new(request)
	r.Context = ctx
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

func (r *request) SessionId() grpc.SessionId {
	return r.sessId
}

func (r *request) NodeId() grpc.NodeId {
	return r.nodeId
}

func (r *request) RequiredServices() []string {
	return nil
}
