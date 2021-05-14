// +build auth

package main

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
)

func init() {
	//addCreateDb()
	//addJoinCommand()
	fmt.Printf("with auth\n")
	mainJob.AddTask(func(j job.Job) (job.Init, job.Run, job.Finalize) {
		run := func(task job.Task) {
			addJoinCommand()
			task.Done()
		}
		return nil, run, nil
	})
}
