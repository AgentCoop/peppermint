package cmd

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
)

func RunCmd() error {
	serviceJob := job.NewJob(nil)
	regServices := runtime.GlobalRegistry().Services()
	for _, desc := range regServices {
		service := desc.Initializer()
		serviceJob.AddTask(service.StartTask)
	}
	<-serviceJob.Run()
	_, err := serviceJob.GetInterruptedBy()
	switch v := err.(type) {
	case error:
		return v
	default:
		return nil
	}
}
