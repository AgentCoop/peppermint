package model

import (
	"github.com/AgentCoop/peppermint/internal"
)

func (d hubDb) FetchById(id internal.NodeId) (*HubJoinedNode, error) {
	db := d.Handle()
	found := &HubJoinedNode{}
	db.Where(&HubJoinedNode{ExternalId: id})
	err := db.First(found).Error
	return found, err
}

func (d hubDb) SaveJoinRequest(id internal.NodeId, encKey []byte) error {
	db := d.Handle()
	node := &HubJoinedNode{
		ExternalId:   id,
		EncKey:       encKey,
	}
	return db.Save(node).Error
}
