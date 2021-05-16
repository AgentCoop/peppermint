package join

import (
	job "github.com/AgentCoop/go-work"
)

func JoinCmd(address string, secret string) job.Job {
	jctx := NewJoinContext(address, secret)
	j :=  job.NewJob(jctx)

	j.AddOneshotTask(jctx.HubClient.ConnectTask)
	j.AddTask(jctx.JoinCmdTask)
	j.AddTask(jctx.JoinHelloTask)
	j.AddTask(jctx.JoinTask)

	return j
}
