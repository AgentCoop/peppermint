package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"

)

type app struct {
	runtime.Runtime
}

func AppInitTest() job.Job {
	app := new(app)
	app.Runtime = runtime.NewRuntime(
		cliparser.NewParser(&cmd.Options),
		&cmd.Options.AppDir,
		"test.db",
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
	app := new(app)
	app.Runtime = runtime.NewRuntime(
		cliparser.NewParser(&cmd.Options),
		&cmd.Options.AppDir,
		dbFilename,
	)
	appJob := job.NewJob(nil)
	appJob.AddOneshotTask(app.InitTask)
	appJob.AddTask(app.ParserTask)
	return appJob
}
