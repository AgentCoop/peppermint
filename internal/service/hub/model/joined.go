package model

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/model"
)

type HubJoinedNode struct {
	model.Model
	JoinAccepted   int
	EncKey         []byte
	E2E_EncEnabled uint
	ExternalId     internal.NodeId `gorm:"type:uint64"`
	Tags           []HubNodeTag    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type HubNodeTag struct {
	model.Model
	Name            string
	HubJoinedNodeID uint
}
