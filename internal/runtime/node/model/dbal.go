package model

import (
	"github.com/AgentCoop/peppermint/internal"
)

func Migrate() {
	//db := runtime.GlobalRegistry().Db().Handle()
	//mig := db.Migrator()
	//job.Logger(logger.Debug)("migrating node tables...")
	//mig.AutoMigrate(tables...)
}

func (n nodeDb) CreateNode(id internal.NodeId, tags []string) (*node, error) {
	db := n.Handle()
	tagModels := make([]NodeTag, len(tags))
	for i, tagName := range tags {
		tagModels[i] = NodeTag{
			Name:   tagName,
		}
	}
	values := &Node{
		ExternalId: id,
		Tags:       tagModels,
	}
	err := db.Create(values).Error
	return &node{values}, err
}

func (n nodeDb) FetchByExternalId(id internal.NodeId) (*Node, error) {
	db := n.Handle()
	first := &Node{}
	err := db.Where(&Node{ExternalId: id}).First(first).Error
	return first, err
}

func (n nodeDb) FetchById(id uint) (*Node, error) {
	db := n.Handle()
	first := &Node{}
	err := db.First(&first, id).Error
	return first, err
}

func (n nodeDb) UpdateNode(nodeId uint, encKey []byte) error {
	db := n.Handle()
	node := &Node{}
	node.ID = nodeId
	vals := &Node{}
	if encKey != nil {
		vals.EncKey = encKey
	}
	err := db.Model(node).Updates(vals).Error
	return err
}
