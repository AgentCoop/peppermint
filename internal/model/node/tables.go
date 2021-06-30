package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/AgentCoop/peppermint/internal/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

type Node struct {
	model.Model
	ExternalId internal.NodeId `gorm:"type:uint64"`
	Tags       []NodeTag       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	EncKey     []byte
	IsSecure   uint
	LbWeight   uint // Load balancer node weight
}

type SecurityPolicy interface {
	E2E_Enabled() bool
}

type NodeTag struct {
	model.Model
	Name   string `gorm:"unique"`
	NodeID uint
}

var tables = []interface{}{
	&Node{}, &NodeTag{},
}

func CreateTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("creating node tables...")
	mig.CreateTable(tables...)
}

func DropTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("dropping node tables...")
	mig.DropTable(tables...)
}
