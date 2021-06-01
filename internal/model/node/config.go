package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/model"
)

type NodeConfig struct {
	model.Model
	Id internal.NodeId `gorm:"type:uint64"`
}
