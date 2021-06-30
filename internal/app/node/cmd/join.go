package cmd

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/client"
	"github.com/AgentCoop/peppermint/internal/service/hub/grpc/client/join"
	"github.com/AgentCoop/peppermint/internal/utils"
)

func JoinCmd(secret string, tags []string, hubAddr string) error {
	joinCtx := join.NewJoinContext(secret, tags)
	hubClient := client.NewClient(hubAddr)

	clientJob := job.NewJob(hubClient)
	clientJob.AddOneshotTask(hubClient.ConnectTask)
	clientJob.AddTask(joinCtx.JoinTask)
	<-clientJob.Run()

	if _, err := clientJob.GetInterruptedBy(); err != nil {
		return utils.Conv_InterfaceToError(err)
	}
	return nil
}
