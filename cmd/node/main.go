package main

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node"
	"github.com/AgentCoop/peppermint/internal/logger"
	//_ "github.com/AgentCoop/peppermint/internal/service/balancer"
	_ "github.com/AgentCoop/peppermint/internal/service/hub"
	//_ "github.com/AgentCoop/peppermint/internal/service/test"
	"github.com/AgentCoop/peppermint/internal/utils"
	"os"
)

const (
	DbFilename = "node.db"
)

func main() {
	appJob := node.NewAppJob()
	<-appJob.Run()

	if _, err := appJob.GetInterruptedBy(); err != nil {
		job.Logger(logger.Error)("%s", utils.Conv_InterfaceToError(err).Error())
		os.Exit(1)
	}
}
