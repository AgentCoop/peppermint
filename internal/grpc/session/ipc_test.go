package session

import (
	"errors"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

type dataBag struct {
	msg string
	num int
}

func TestCommunicator_DataPingPong(t *testing.T) {
	comm := newIpc(Sequential)
	j := comm.Job()
	ping, pong := "ping", "pong"
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			expected := ping
			c := j.GetValue().(grpc.GrpcServiceLayersIpc)
			data := c.Svc_Recv(0)
			task.AssertNotNil(data)
			v := data.(*dataBag)
			if v.msg != expected {
				t.Fatalf("expected %s, got %s", expected, v.msg)
			}
			v.msg = pong
			c.Svc_Send(0, data)
			task.Done()
		}
		return nil, run, nil
	})
	data := &dataBag{}
	go func() {
		data.msg = "ping"
		expected := "pong"
		comm.Grpc_Send(0, data)
		rxData := comm.Grpc_Recv(0)
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
	comm := newIpc(Sequential)
	j := comm.Job()
	svcErr := errors.New("service error")
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			c := j.GetValue().(grpc.GrpcServiceLayersIpc)
			_ = c.Svc_Recv(0)
			task.Assert(svcErr) // trigger some error
			task.Done()
		}
		return nil, run, nil
	})
	j.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			time.Sleep(50 * time.Millisecond)
			c := j.GetValue().(grpc.GrpcServiceLayersIpc)
			data := c.Svc_Recv(1)
			task.AssertNotNil(data)
			task.Done()
		}
		return nil, run, nil
	})
	data := &dataBag{}
	go func() {
		comm.Grpc_Send(0, data)
		rxData := comm.Grpc_Recv(0)
		switch v := rxData.(type) {
		case error:
			st := status.Convert(rxData.(error))
			if st.Message() != svcErr.Error() {
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
	tick int
	jobsize int
	usechan int
}

func (c *cruncher) StreamableNumCruncher(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		comm := j.GetValue().(grpc.GrpcServiceLayersIpc)
		data := comm.Svc_Recv(c.usechan) // service <- grpc
		task.AssertNotNil(data)
		num := data.(*dataBag).num
		data.(*dataBag).num = num * num
		comm.Svc_Send(c.usechan, data) // service -> grpc
		c.tick++
		if c.tick == c.jobsize {
			task.Done()
		} else {
			task.Tick()
		}
	}
	return nil, run, nil
}

func TestCommunicator_Streamable(t *testing.T) {
	comm := newIpc(Sequential)
	N := 20
	var recvx int
	j := comm.Job()
	crunch := &cruncher{0, N, 0}
	j.AddTask(crunch.StreamableNumCruncher)
	for i := 0; i < N; i++ {
		num := i
		go func() {
			data := &dataBag{}
			data.num = num
			comm.Grpc_SendStreamable(crunch.usechan, data)
			rxData := comm.Grpc_Recv(crunch.usechan)
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
	case crunch.tick != N:
		t.Fatalf("expected %d to be crunched, got %d", N, crunch.tick)
	}
}

func TestCommunicator_OutOfOrder(t *testing.T) {
	comm := newIpc(OutOfOrder)
	j := comm.Job()
	crunch := &cruncher{0, 1, 0}
	crunch2 := &cruncher{0, 1, 1}
	j.AddTask(crunch.StreamableNumCruncher)
	j.AddTask(crunch2.StreamableNumCruncher)
	// Out of order call
	go func() {
		data := &dataBag{}
		data.num = 3
		comm.Grpc_Send(1, data)
		rxData := comm.Grpc_Recv(1)
		if rxData.(*dataBag).num != 9 {
			t.Fatalf("expected %d, got %v", 9, rxData)
		}
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		data := &dataBag{}
		data.num = 2
		comm.Grpc_Send(0, data)
		rxData := comm.Grpc_Recv(0)
		if rxData.(*dataBag).num != 4 {
			t.Fatalf("expected %d, got %v", 4, rxData)
		}
	}()
	<-j.Run()
	switch {
	case j.GetState() != job.Done:
		t.Fatalf("expected job state %s, got %s", job.Done, j.GetState())
	case crunch.tick != 1 || crunch2.tick != 1:
		t.Fatalf("expected %d to be crunched, got %d %d", 2, crunch.tick, crunch2.tick)
	}
}