package hub

import (
	job "github.com/AgentCoop/go-work"
)

type service struct {

}

func NewService() *service {
	s := &service{}
	return s
}

func (s *service) CreateDbTask() job.JobTask {
	return func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			task.Done()
		}
		return nil, run, nil
	}
}

func (s *service) Start() {

}
