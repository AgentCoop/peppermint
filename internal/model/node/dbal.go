package node

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func CreateTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	mig.CreateTable(&Node{}, &NodeTag{})
}

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
	err := db.Create(node).Error
	return err
}

func UpdateNodeEncKey(encKey []byte) error {
	db := runtime.GlobalRegistry().Db().Handle()
	node := Node{}
	db.First(&node)
	node.EncKey = encKey
	err := db.Save(&node).Error
	return err
}
