package model

import (
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func FetchById(id internal.NodeId) (*HubJoinedNode, error) {
	db := runtime.GlobalRegistry().Db().Handle()
	found := &HubJoinedNode{}
	db.Where(&HubJoinedNode{ExternalId: id})
	err := db.First(found).Error
	return found, err
}

func SaveJoinRequest(id internal.NodeId, encKey []byte) error {
	db := runtime.GlobalRegistry().Db().Handle()
	node := &HubJoinedNode{
		ExternalId: id,
		JoinAccepted: 0,
		EncKey: encKey,
	}
	return db.Save(node).Error
}

func AcceptJoin(id internal.NodeId, encKey []byte) error {
	db := runtime.GlobalRegistry().Db().Handle()
	db.Model(&HubJoinedNode{})
	db.Where(&HubJoinedNode{ExternalId: id})
	db.Updates(HubJoinedNode{
		EncKey: encKey,
	})
	return db.Error
}

func RejectJoin(id internal.NodeId) error {
	db := runtime.GlobalRegistry().Db().Handle()
	db.Where(&HubJoinedNode{ExternalId: id, JoinAccepted: 0})
	db.Delete(&HubJoinedNode{})
	return db.Error
}
