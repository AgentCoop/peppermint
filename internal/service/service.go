package service

import (
	job "github.com/AgentCoop/go-work"
)

type Service interface {
	CreateDbTask() job.JobTask
	StartTask() job.JobTask
}
