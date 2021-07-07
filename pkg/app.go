package pkg

import (
	job "github.com/AgentCoop/go-work"
)

type App interface {
	RootDir() string
	Db() Db
	WithDb(Db)
	Job() job.Job
}

type AppNode interface {
	App
	Node() Node
	ServiceDb() Db
	InitServiceDb(node Node) error
}
