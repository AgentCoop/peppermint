package grpc

import (
	"github.com/AgentCoop/peppermint/internal/api/peppermint/service"
)

type NodeId uint64
type SessionId uint64

const (
	META_FIELD_NODE_ID = "gs-node-id"
	META_FIELD_SESSION_ID = "gs-session-id"
)

type customErr struct {
	statusCode service.StatusCode
	text string
}

func NewCustomError(code service.StatusCode, text string) {

}