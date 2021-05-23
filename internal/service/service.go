package service

import (
	job "github.com/AgentCoop/go-work"
)

type Service interface {
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
}