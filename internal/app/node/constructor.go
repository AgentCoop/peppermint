package node

import (
	job "github.com/AgentCoop/go-work"
	app "github.com/AgentCoop/peppermint/internal/app"
	"github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"os"
)

func NewApp(profile app.AppProfile) *appNode {
	appNode := new(appNode)
	appNode.App = app.NewApp(profile, &cmd.Options)
	appNode.Job().AddOneshotTask(appNode.InitTask)
	runtime.GlobalRegistry().SetApp(appNode)
	return appNode
}

func NewTestAppJob() job.Job {
	profile, _ := app.AppProfiles[app.TEST]
	os.Remove(profile.DbFilename)

	nodeApp := NewApp(profile)
	appJob := nodeApp.Job()
	appJob.AddTask(nodeApp.SetupTestDbTask)
	<-appJob.Run()

	nodeApp = NewApp(profile)
	appJob = nodeApp.Job()
	appJob.AddTask(nodeApp.ParserTask)
	return appJob
}

func NewAppJob() job.Job {
	prof := app.ProfileFromEnv()
	nodeApp := NewApp(prof)
	appJob := nodeApp.Job()
	appJob.AddTask(nodeApp.ParserTask)
	return appJob
}
