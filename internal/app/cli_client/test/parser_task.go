package test

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/service/test/grpc/client"
)

func (app *appTest) ParserTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		rt := runtime.GlobalRegistry().Runtime()
		parser := rt.CliParser()
		cmdName, _ := parser.CurrentCmd()
		o := &options
		callParams := &callParams{
			token:          o.Token,
			repeat:         o.CallRepeatCount,
			workers:        o.CallWorkersCount,
			rsDelay:        o.RsDelay,
			rsDelayJitter:  o.RsDelayJitter,
			rqBulkMin:      o.RqBulkDataMin,
			rqBulkMax:      o.RqBulkDataMax,
			callFailProbab: o.CallFailureProbab,
		}
		// Normalize params
		if callParams.repeat == 0 {
			callParams.repeat = 1
		}
		if callParams.workers == 0 {
			callParams.workers = 1
		}
		app.callParams = callParams

		cc := client.NewClient(options.Host)
		if options.Port > 0 {
			cc.WithTargetPort(options.Port)
		}

		// Pass a command to execute to the executor task
		switch cmdName {
		case CMD_NAME_SINGLE:
			app.execCmdChan <- cc.Single
		default:
			task.Assert(fmt.Errorf("unknown command %s", cmdName))
		}
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}
