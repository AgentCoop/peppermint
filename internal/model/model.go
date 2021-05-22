package model

import "github.com/AgentCoop/peppermint/internal/db"

type Model interface {
	Create() db.PrimaryKey
	Read(key db.PrimaryKey) Model
	Update() Model
	Delete()
	SqlDDLStatement() string
}
