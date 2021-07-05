package model

import "github.com/AgentCoop/peppermint/pkg"

type nodeDb struct {
	pkg.Db
}

func NewDb(db pkg.Db) nodeDb {
	return nodeDb{db}
}

func (n nodeDb) NewNode(id uint) (*node, error) {
	m, err := n.FetchById(id)
	if err != nil {
		return nil, err
	}
	node := &node{m}
	return node, nil
}
