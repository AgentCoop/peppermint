package grpc_test

import (
	"github.com/AgentCoop/go-work"
	server2 "github.com/AgentCoop/peppermint/internal/service/hub/grpc/server"
	cmd "github.com/AgentCoop/peppermint/internal/service/hub/service/client/join"
	"time"

	"testing"
)

var (
	serverAddr = "localhost:9911"
)

func TestJoinHello(t *testing.T) {
	serverJob := job.NewJob(t)
	server := server2.NewServer(serverAddr)
	serverJob.AddTask(server.StartTask)
	//j.AddTask(createClient)
	serverJob.Run()

	time.Sleep(100 * time.Millisecond)
	clientJob := cmd.JoinCmd(serverAddr, "secretword")
	<-clientJob.Run()

	_, err := clientJob.GetInterruptedBy()
	if err != nil {
		t.Error(err)
	}
}
