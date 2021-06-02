package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/node"
	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
)

type app struct {
	runtime.Runtime
	appDir *string
	dbFilename string
}

func AppInitTest() job.Job {
	app := &app{
		appDir: &cmd.Options.AppDir,
		dbFilename: "test.db",
	}
	app.Runtime = runtime.NewRuntime(
		node.NewConfigurator(),
		cliparser.NewParser(&cmd.Options),
	)

	dbJob := job.NewJob(nil)
	dbJob.AddOneshotTask(app.InitTask)
	dbJob.AddTask(app.SetupTestDbTask)
	<-dbJob.Run()

	appJob := job.NewJob(nil)
	appJob.AddOneshotTask(app.InitTask)
	appJob.AddTask(app.ParserTask)

	return appJob
}

func AppInit(dbFilename string) job.Job {
	app := &app{
		appDir: &cmd.Options.AppDir,
		dbFilename: dbFilename,
	}
	app.Runtime = runtime.NewRuntime(
		node.NewConfigurator(),
		cliparser.NewParser(&cmd.Options),
	)
	appJob := job.NewJob(nil)
	appJob.AddOneshotTask(app.InitTask)
	appJob.AddTask(app.ParserTask)
	return appJob
}
