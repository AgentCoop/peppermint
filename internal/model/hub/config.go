package hub

import "github.com/AgentCoop/peppermint/internal/model"

type HubConfig struct {
	model.Model
	Port int `gorm:"default:12000"`
	Address string `gorm:"default:localhost"`
}
