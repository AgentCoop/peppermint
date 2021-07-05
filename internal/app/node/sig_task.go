package node

import (
	job "github.com/AgentCoop/go-work"
	"syscall"
)

func (appNode *appNode) SigTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		sig := <-appNode.sigChan
		switch sig {
		case syscall.SIGHUP:
			appNode.reloadServiceConfig()
		default:
			task.Done()
			return
		}
		task.Tick()
	}
	fin := func(task job.Task) {
		close(appNode.sigChan)
	}
	return nil, run, fin
}

