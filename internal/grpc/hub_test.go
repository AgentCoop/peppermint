package grpc_test

import (
	"github.com/AgentCoop/go-work"
	c "github.com/AgentCoop/peppermint/internal/grpc/client/hub"
	h "github.com/AgentCoop/peppermint/internal/grpc/server/hub"
	"time"

	"testing"
)

func createClient(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
	}
	run := func(task job.Task) {
		time.Sleep(100 * time.Millisecond)
		client := c.NewClient("localhost:9000", task)
		client.Connect()
		client.JoinHello([]byte("hello"))
		task.FinishJob()
	}
	return init, run, nil
}

func TestJoinHello(t *testing.T) {
	j := job.NewJob(t)
	server := h.NewServer()
	j.AddTask(server.StartServerTask)
	j.AddTask(createClient)
	<-j.Run()
	_, err := j.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
