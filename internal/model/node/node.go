package node

import (
	"github.com/AgentCoop/peppermint/internal/db"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/model"
	"net"
)

type node struct {
	Id db.PrimaryKey
	EncKey []byte
	Tags []string
}

type Node interface {
	model.Model
	Id() grpc.NodeId
	ServiceAddressByName(string) net.Addr
}
