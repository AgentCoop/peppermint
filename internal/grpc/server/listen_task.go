package server

import (
	job "github.com/AgentCoop/go-work"
	"net"
)

func (s *baseServer) ListenTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		lis, err := net.Listen(s.address.Network(), s.address.String())
		task.Assert(err)
		s.lis = lis
	}
	run := func (task job.Task) {
		s.handle.Serve(s.lis)
		task.Done()
	}
	return init, run, nil
}
