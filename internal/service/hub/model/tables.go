package model

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/logger"
	"gorm.io/gorm"
)

type HubConfig struct {
	gorm.Model
	Port    int    `gorm:"default:12000"`
	Address string `gorm:"default:localhost"`
	Secret  string `gorm:"default:secret"`
}

type HubJoinedNode struct {
	gorm.Model
	Status         uint
	EncKey         []byte
	E2E_EncEnabled uint
	ExternalId     internal.NodeId `gorm:"type:uint64"`
	Tags           []HubNodeTag    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type HubNodeTag struct {
	gorm.Model
	Name            string
	HubJoinedNodeID uint
}

var (
	tables = []interface{}{
		&HubConfig{}, &HubJoinedNode{}, &HubNodeTag{},
	}
)

func (h hubDb) CreateTables() {
	db := h.Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("creating Hub tables...")
	mig.CreateTable(tables...)
}

func (h hubDb) DropTables() {
	db := h.Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("dropping Hub tables...")
	mig.DropTable(tables...)
}
