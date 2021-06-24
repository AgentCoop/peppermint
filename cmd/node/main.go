package main

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node"
	"github.com/AgentCoop/peppermint/internal/logger"
	_ "github.com/AgentCoop/peppermint/internal/plugin/hub"
	_ "github.com/AgentCoop/peppermint/internal/plugin/balancer"
	"github.com/AgentCoop/peppermint/internal/utils"
	"os"
)

const (
	DbFilename = "node.db"
)

func main() {
	appJob := node.AppInit(DbFilename)
	<-appJob.Run()

	if _, err := appJob.GetInterruptedBy(); err != nil {
		job.Logger(logger.Error)("%s", utils.Conv_InterfaceToError(err).Error())
		os.Exit(1)
	}
}
