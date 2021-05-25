package main

import (
	"github.com/AgentCoop/peppermint/internal/app/node"
	_ "github.com/AgentCoop/peppermint/internal/service/hub"
	_ "github.com/AgentCoop/peppermint/internal/service/webproxy"
	"os"
)

const (
	DbFilename = "node.db"
)

func main() {
	appJob := node.AppJob(DbFilename)
	<-appJob.Run()

	_, err := appJob.GetInterruptedBy()
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
