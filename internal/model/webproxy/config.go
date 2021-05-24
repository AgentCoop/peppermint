package webproxy

import "github.com/AgentCoop/peppermint/internal/model"

type WebProxy struct {
	model.Model
	Port int
}
