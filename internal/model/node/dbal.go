package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func CreateNode(id internal.NodeId, tags []string) error {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	mig.DropTable(Tables...)
	mig.CreateTable(Tables...)

	tagModels := make([]NodeTag, len(tags))
	for i, tagName := range tags {
		tagModels[i] = NodeTag{
			Name:   tagName,
		}
	}

	node := &Node{
		ExternalId: id,
		Tags:       tagModels,
	}
	return db.Create(node).Error
}