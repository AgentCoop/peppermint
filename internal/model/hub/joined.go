package hub

import (
	"github.com/AgentCoop/peppermint/internal/model"
)

type HubJoinedNode struct {
	model.Model
	EncKey []byte
	NodeId uint64 `gorm:"type:uint64"`
	Tags []HubNodeTag `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type HubNodeTag struct {
	model.Model
	Name string
	HubJoinedNodeID uint
}
