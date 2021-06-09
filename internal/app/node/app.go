package node

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/app/node/cmd"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
	"github.com/AgentCoop/peppermint/internal/runtime/node"
	"os"
)

const (
	DEV_DB_NAME = "test.db"
	PROD_DB_NAME = "node.db"
)

type app struct {
	runtime.Runtime
	appDir     *string
	dbFilename string
	dbPathname string
}

func NewApp(dbFilename string) *app {
	app := &app{
		appDir:     &cmd.Options.AppDir,
		dbFilename: dbFilename,
	}
	app.Runtime = runtime.NewRuntime(
		node.NewNodeManager(),
		node.NewConfigurator(),
		cliparser.NewParser(&cmd.Options),
	)
	runtime.GlobalRegistry().SetRuntime(app)
	return app
}

func AppInitTest() job.Job {
	os.Remove(DEV_DB_NAME)
	app := NewApp(DEV_DB_NAME)

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
	app := NewApp(dbFilename)
	appJob := job.NewJob(nil)
	appJob.AddOneshotTask(app.InitTask)
	appJob.AddTask(app.ParserTask)
	return appJob
}
