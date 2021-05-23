package runtime

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
)

func (r *runtime) NodeTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		fmt.Printf("run node task\n")
		task.Done()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

