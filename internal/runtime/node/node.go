package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
)

type cfg struct {
	nodeModel node.Node
}

func NewConfigurator() *cfg {
	cfg := new(cfg)
	return cfg
}

func (c *cfg) Fetch() error {
	db := runtime.GlobalRegistry().Db().Handle()
	return db.First(&c.nodeModel).Error
}

func (c *cfg) MergeCliOptions(parser deps.CliParser) {

}

func (c *cfg) NodeId() internal.NodeId {
	return c.nodeModel.ExternalId
}

func (c *cfg) Tags() []string {
	tags := make([]string, len(c.nodeModel.Tags))
	for i, v := range c.nodeModel.Tags {
		tags[i] = v.Name
	}
	return tags
}

func (c *cfg) EncKey() []byte {
	return c.nodeModel.EncKey
}

func (c *cfg) IsSecure() bool {
	switch {
	case c.nodeModel.IsSecure > 0:
		return true
	default:
		return false
	}
}
