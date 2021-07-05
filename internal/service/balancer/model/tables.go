package model

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/AgentCoop/peppermint/internal/runtime/node/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

type BalancerConfig struct {
	model.Model
	Port          int                     `gorm:"default:12001"`
	Address       string                  `gorm:"default:localhost"`
	DefaultAlgo   int                     `gorm:"default:0"` // Random
	PreferredAlgo []BalancerPreferredAlgo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BalancerPreferredAlgo struct {
	model.Model
	ServiceName      string `gorm:"not null"`
	Algo             int    `gorm:"default:0"` // Random
	BalancerConfigID uint
}

var tables = []interface{}{
	&BalancerConfig{}, &BalancerPreferredAlgo{},
}

func CreateTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("creating Balancer tables...")
	mig.CreateTable(tables...)
}

func DropTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("dropping Balancer tables...")
	mig.DropTable(tables...)
}
