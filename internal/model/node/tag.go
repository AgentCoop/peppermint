package node

import (
	"github.com/AgentCoop/peppermint/internal/model"
)

type NodeTag struct {
	model.Model
	Name string `gorm:"unique"`
	NodeID uint
}
