package runtime

import (
	//	"fmt"
	"fmt"
	job "github.com/AgentCoop/go-work"
	//	"github.com/AgentCoop/peppermint/internal/runtime/cliparser"
)

func (r *runtime) InitTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		fmt.Printf("parse\n")
		err := r.CliParser.Run()
		task.Assert(err)
		fmt.Printf("init runtime\n")
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

