package main

import (
	"github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
	_ "github.com/AgentCoop/peppermint/internal/service/hub"
	_ "github.com/AgentCoop/peppermint/internal/service/webproxy"
	"os"
)

const (
	DbFilename = "node.db"
)

type app struct {
	runtime.Runtime
}

func addServices(j job.Job, app *app) {
	//for _, service := range app.Services() {
	//	j.AddTask(service.StartTask)
	//}
}

func main() {
	app := new(app)
	app.Runtime = runtime.NewRuntime(
		cliparser.NewParser(&options),
		DbFilename,
	)
	appJob := job.NewJob(nil)
	appJob.AddOneshotTask(app.InitTask)
	appJob.AddTask(app.NodeTask)
	<-appJob.Run()

	_, err := appJob.GetInterruptedBy()
	if err != nil {
		panic(err)
	}
	os.Exit(0)
}
