package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func CreateTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.DbKey)("creating node tables...")
	mig.CreateTable(tables...)
}

func DropTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.DbKey)("dropping node tables...")
	mig.DropTable(tables...)
}

func CreateNode(id internal.NodeId, tags []string) error {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	mig.DropTable(tables...)
	mig.CreateTable(tables...)

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
