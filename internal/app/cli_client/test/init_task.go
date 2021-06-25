package test

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app"
	"github.com/AgentCoop/peppermint/internal/service/test/grpc/client"
	"net"
)

func (appTest *appTest) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		app.AppInit(appTest, task)
		addr, err := net.ResolveTCPAddr("tcp", Options.Service)
		task.Assert(err)
		cc := client.NewClient(addr)
		j.SetValue(cc)
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}
