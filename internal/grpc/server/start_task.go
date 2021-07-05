package server

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/utils/str"
	"net"
)

func (srv *baseServer) StartTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		lis, err := net.Listen(srv.address.Network(), srv.address.String())
		task.Assert(err)
		srv.lis = lis
	}
	run := func (task job.Task) {
		if srv.logger != nil {
			srvName := str.SplitAndLast(srv.fullName, ".")
			srv.logger("%s started to serve requests on %s", srvName, srv.address.String())
		}
		srv.handle.Serve(srv.lis)
		task.Done()
	}
	return init, run, nil
}
