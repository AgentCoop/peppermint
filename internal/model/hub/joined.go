package hub

import "github.com/AgentCoop/peppermint/internal/model"

// Holds joined nodes
type JoinedNode struct {
	model.Model
	EncKey string
}
