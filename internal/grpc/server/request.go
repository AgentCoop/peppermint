package server

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/metadata"
	"context"
)

type request struct {
	context.Context
	sessId i.SessionId
	nodeId i.NodeId
}

type RequestHeader interface {
	SessionId() i.SessionId
	NodeId() i.NodeId
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
	//
	var vals []string
	vals = md.Get(grpc.META_FIELD_NODE_ID)
	if len(vals) > 1 {
		r.nodeId = i.NodeId(utils.Hex2int(vals[0]))
	}
	return r
}

func (r *request) SessionId() i.SessionId {
	return r.sessId
}

func (r *request) NodeId() i.NodeId {
	return r.nodeId
}

func (r *request) RequiredServices() []string {
	return nil
}

type reqHeader struct {

}

func NewRequestHeader(md metadata.MD) *reqHeader {
	r := new(reqHeader)
	return r
}

