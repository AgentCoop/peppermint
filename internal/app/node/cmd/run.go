package cmd

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func RunCmd() error {
	svcJob := job.NewJob(nil)
	rt := runtime.GlobalRegistry().Runtime()
	for _, svc := range rt.Services() {
		svcJob.AddTask(svc.InitTask)
		uds, tcp := svc.IpcServer(), svc.Server()
		if uds != nil {
			svcJob.AddTask(uds.StartTask)
		}
		if tcp != nil {
			svcJob.AddTask(tcp.StartTask)
		}
	}
	<-svcJob.Run()
	_, err := svcJob.GetInterruptedBy()
	switch v := err.(type) {
	case error:
		return v
	default:
		return nil
	}
}
