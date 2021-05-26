package server_test

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/utils"
	"testing"
)

type dataBag struct {
	msg string
}

func TestCommunicator_DataPingPong(t *testing.T) {
	comm := server.NewCommunicator()
	j := utils.DefaultGrpcJob(comm)
	ping, pong := "ping", "pong"
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			expected := ping
			c := j.GetValue().(runtime.GrpcServiceCommunicator)
			data := c.ServiceRx(0)
			task.AssertNotNil(data)
			v := data.(*dataBag)
			if v.msg != expected {
				t.Fatalf("expected %s, got %s", expected, v.msg)
			}
			v.msg = pong
			c.ServiceTx(0, data)
			task.Done()
		}
		return nil, run, nil
	})
	data := &dataBag{}
	go func() {
		data.msg = "ping"
		expected := "pong"
		comm.GrpcTx(0, data)
		rxData := <-comm.GrpcRx(0)
		switch rxData.(type) {
		case error:
			t.Fatal(rxData)
		default:
			if rxData.(*dataBag).msg != expected {
				t.Fatalf("expected %s, got %s", expected, data.msg)
			}
		}
	}()
	<-j.Run()
}

