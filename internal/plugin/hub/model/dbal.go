package model

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

var (
	tables = []interface{} {
		&HubConfig{}, &HubJoinedNode{}, &HubNodeTag{},
	}
)

func CreateTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("creating Hub tables...")
	mig.CreateTable(tables...)
}

func DropTables() {
	db := runtime.GlobalRegistry().Db().Handle()
	mig := db.Migrator()
	job.Logger(logger.Debug)("dropping Hub tables...")
	mig.DropTable(tables...)
}

func FetchById(id internal.NodeId) (*HubJoinedNode, error) {
	db := runtime.GlobalRegistry().Db().Handle()
	found := &HubJoinedNode{}
	db.Where(&HubJoinedNode{ExternalId: id})
	err := db.First(found).Error
	return found, err
}

func HasJoined(id internal.NodeId) bool {
	db := runtime.GlobalRegistry().Db().Handle()
	node := HubJoinedNode{ExternalId: id, JoinAccepted: 1}
	result := db.Where(&node).First(&node)
	return result.RowsAffected > 0
}

func SaveJoinRequest(id internal.NodeId, encKey []byte) error {
	db := runtime.GlobalRegistry().Db().Handle()
	node := &HubJoinedNode{
		ExternalId:   id,
		JoinAccepted: 0,
		EncKey:       encKey,
	}
	return db.Save(node).Error
}

func AcceptJoin(id internal.NodeId) error {
	db := runtime.GlobalRegistry().Db().Handle()
	db.Model(HubJoinedNode{}).Where("external_id = ?", id).Updates(&HubJoinedNode{
		JoinAccepted: 1,
	})
	return db.Error
}

func RejectJoin(id internal.NodeId) error {
	db := runtime.GlobalRegistry().Db().Handle()
	db.Where(&HubJoinedNode{ExternalId: id, JoinAccepted: 0})
	db.Delete(&HubJoinedNode{})
	return db.Error
}
