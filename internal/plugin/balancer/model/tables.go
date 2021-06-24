package model

import "github.com/AgentCoop/peppermint/internal/model"

type BalancerConfig struct {
	model.Model
	Port          int                     `gorm:"default:12001"`
	Address       string                  `gorm:"default:localhost"`
	DefaultAlgo   int                     `gorm:"default:0"` // Random
	PreferredAlgo []BalancerPreferredAlgo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type BalancerPreferredAlgo struct {
	model.Model
	ServiceName    string `gorm:"not null"`
	Algo           int    `gorm:"default:0"` // Random
	BalancerConfig uint
}
