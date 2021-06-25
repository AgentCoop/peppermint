package pkg

import job "github.com/AgentCoop/go-work"

type App interface {
	WithDb()
	Job() job.Job
	RootDir() string
}
