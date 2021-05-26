package server_test

import (
	"errors"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/server"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"google.golang.org/grpc/status"
	"testing"
)

type dataBag struct {
	msg string
	num int
}

func TestCommunicator_DataPingPong(t *testing.T) {
	comm := server.NewCommunicator()
	j := comm.Job()
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
		rxData := comm.GrpcRx(0)
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

func TestCommunicator_ErrorPropagation(t *testing.T) {
	comm := server.NewCommunicator()
	j := comm.Job()
	svcErr := errors.New("service error")
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			c := j.GetValue().(runtime.GrpcServiceCommunicator)
			_ = c.ServiceRx(0)
			task.Assert(svcErr) // trigger some error
			task.Done()
		}
		return nil, run, nil
	})
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			c := j.GetValue().(runtime.GrpcServiceCommunicator)
			data := c.ServiceRx(1)
			task.AssertNotNil(data)
			task.Done()
		}
		return nil, run, nil
	})
	data := &dataBag{}
	go func() {
		comm.GrpcTx(0, data)
		rxData := comm.GrpcRx(0)
		switch v := rxData.(type) {
		case *status.Status:
			if rxData.(*status.Status).Message() != svcErr.Error() {
				t.Fatalf("expected %v", svcErr)
			}
		default:
			_ = v
			t.Fatalf("expected gRPC error, got %v", rxData)
		}
	}()
	<-j.Run()
	if j.GetState() != job.Cancelled {
		t.Fatalf("expected job state %s, got %s", job.Cancelled, j.GetState())
	}
}

type cruncher struct {
	val int
}

func (c *cruncher) StreamableNumCruncher(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		comm := j.GetValue().(runtime.GrpcServiceCommunicator)
		data := comm.ServiceRx(0) // service <- grpc
		task.AssertNotNil(data)
		num := data.(*dataBag).num
		data.(*dataBag).num = num * num
		comm.ServiceTx(0, data) // service -> grpc
		if c.val == 19 {
			task.Done()
		} else {
			c.val++
			task.Tick()
		}
	}
	return nil, run, nil
}

func TestCommunicator_Streamable(t *testing.T) {
	comm := server.NewCommunicator()
	count := 20
	var recvx int
	j := comm.Job()
	crunch := new(cruncher)
	j.AddTask(crunch.StreamableNumCruncher)
	for i := 0; i < count; i++ {
		num := i
		go func() {
			data := &dataBag{}
			data.num = num
			comm.GrpcTxStreamable(0, data)
			rxData := comm.GrpcRx(0)
			recvx++
			if rxData.(*dataBag).num != num * num {
				t.Fatalf("expected %d, got %v", num* num, rxData)
			}
		}()
	}
	<-j.Run()
	switch {
	case j.GetState() != job.Done:
		t.Fatalf("expected job state %s, got %s", job.Done, j.GetState())
	case crunch.val != count - 1:
		t.Fatalf("expected %d to be crunched, got %d", count, crunch.val)
	}
}
