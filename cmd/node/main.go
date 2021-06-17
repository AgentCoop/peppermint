package main

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node"
	"github.com/AgentCoop/peppermint/internal/logger"
	_ "github.com/AgentCoop/peppermint/internal/plugin/hub"
	//_ "github.com/AgentCoop/peppermint/internal/plugin/webproxy"
	"os"
)

const (
	DbFilename = "node.db"
)

func main() {
	appJob := node.AppInit(DbFilename)
	<-appJob.Run()

	_, err := appJob.GetInterruptedBy()
	if err != nil {
		panic(err)
	//	os.Exit(0)
	}
	var errText string
	switch v := err.(type) {
	case error:
		errText = v.Error()
	case string:
		errText = v
	default:
		errText = "unknown"
	}
	job.Logger(logger.Error)("critical error %s", errText)
	os.Exit(-1)
}
