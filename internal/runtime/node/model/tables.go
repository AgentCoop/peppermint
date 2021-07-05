package model

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/logger"
	"gorm.io/gorm"
)

type Node struct {
	gorm.Model
	ExternalId internal.NodeId `gorm:"type:uint64"`
	AppId      string          `gorm":"unique"`
	Hostname   string
	Port       uint16
	EncKey     []byte
	EncEnabled uint
	LbWeight   int           // Load balancer node weight
	Tags       []NodeTag     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Services   []NodeService `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type NodeService struct {
	gorm.Model
	Name   string
	Port   uint16
	NodeID uint
}

type NodeTag struct {
	gorm.Model
	Name   string `gorm:"unique"`
	NodeID uint
}

type SecurityPolicy interface {
	E2E_Enabled() bool
}

var tables = []interface{}{
	&Node{}, &NodeTag{}, &NodeService{},
}

func (n nodeDb) CreateTables() {
	db := n.Handle()
	job.Logger(logger.Debug)("creating node tables...")
	mig := db.Migrator()
	mig.CreateTable(tables...)
}

func (n nodeDb) DropTables() {
	db := n.Handle()
	job.Logger(logger.Debug)("dropping node tables...")
	mig := db.Migrator()
	mig.DropTable(tables...)
}
