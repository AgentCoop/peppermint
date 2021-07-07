package app

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/logger"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
	"github.com/AgentCoop/peppermint/internal/runtime/node"
	"os"
)

func NewApp(opts interface{}) *app {
	var cwd string
	var cwdErr error
	app := &app{}
	app.appJob = job.NewJob(nil)
	cwd = os.Getenv(ENV_ROOT)
	if len(cwd) == 0 {
		cwd, cwdErr = os.Getwd()
		if cwdErr != nil {
			job.Logger(logger.Error)(cwdErr.Error())
			os.Exit(1)
		}
	}
	app.appDir = cwd
	rt := runtime.NewRuntime(
		node.NewNodeManager(),
		cliparser.NewParser(opts),
	)
	runtime.GlobalRegistry().SetRuntime(rt)
	return app
}
