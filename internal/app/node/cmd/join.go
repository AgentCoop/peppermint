package cmd

import (
	"errors"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/grpc/client/join"
	"net"
)

func JoinCmd(secret string, tags []string, hubAddr string) error {
	addr, err := net.ResolveTCPAddr("tcp", hubAddr)
	if err != nil { return err }

	joinCtx := join.NewJoinContext(secret, tags)
	hubClient := client.NewClient(addr)

	clientJob := job.NewJob(hubClient)
	clientJob.AddOneshotTask(hubClient.ConnectTask)
	clientJob.AddTask(joinCtx.JoinTask)
	<-clientJob.Run()

	_, jobErr := clientJob.GetInterruptedBy()
	if jobErr != nil { return errors.New("@todo") }

	return nil
}
