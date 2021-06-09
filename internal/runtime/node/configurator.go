package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
)

func (c *cfg) Fetch() error {
	db := runtime.GlobalRegistry().Db().Handle()
	db.FirstOrCreate(&c.nodeModel)
	c.e2e_EncEnabled = c.nodeModel.IsSecure > 0
	return nil
}

func (c *cfg) MergeCliOptions(parser deps.CliParser) {

}

func (c *cfg) ExternalId() internal.NodeId {
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

func (c *cfg) E2E_EncryptionEnabled() bool {
	return c.e2e_EncEnabled
}
