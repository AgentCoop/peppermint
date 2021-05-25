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

func AppJob(dbFilename string) job.Job {
	app := new(app)
	app.Runtime = runtime.NewRuntime(
		cliparser.NewParser(&cmd.Options),
		dbFilename,
	)
	appJob := job.NewJob(nil)
	appJob.AddOneshotTask(app.InitTask)
	appJob.AddTask(app.ParserTask)
	return appJob
}
