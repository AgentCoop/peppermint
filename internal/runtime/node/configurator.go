package node

import (
	"github.com/AgentCoop/peppermint/internal"
	model "github.com/AgentCoop/peppermint/internal/model/node"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func (n *node) Fetch() error {
	db := runtime.GlobalRegistry().Db().Handle()
	nodeModel := model.Node{}
	err := db.First(&nodeModel).Error
	if err != nil { return err }
	n.encKey = nodeModel.EncKey
	n.externalId = nodeModel.ExternalId
	// Tags
	tags := make([]string, len(nodeModel.Tags))
	for i, v := range nodeModel.Tags {
		tags[i] = v.Name
	}
	n.tags = tags
	n.encEnabled = nodeModel.IsSecure > 0
	return nil
}

func (n *node) Refresh() error {
	return n.Fetch()
}

func (c *node) MergeCliOptions(parser runtime.CliParser) {

}

func (c *node) ExternalId() internal.NodeId {
	return c.externalId
}

func (c *node) Tags() []string {
	return c.tags
}

func (c *node) EncKey() []byte {
	return c.encKey
}

func (c *node) E2E_EncryptionEnabled() bool {
	return c.encEnabled
}
