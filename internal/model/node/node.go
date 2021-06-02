package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/model"
)

type Node struct {
	model.Model
	ExternalId internal.NodeId `gorm:"type:uint64"`
	Tags []NodeTag `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
