package join

import (
	job "github.com/AgentCoop/go-work"
	"net"
)

func JoinCmd(address string, secret string) job.Job {
	addr, _ := net.ResolveTCPAddr("tcp", address)
	jctx := NewJoinContext(addr, secret)
	j :=  job.NewJob(jctx)

	j.AddOneshotTask(jctx.HubClient.ConnectTask)
	j.AddTask(jctx.JoinCmdTask)
	j.AddTask(jctx.JoinHelloTask)
	j.AddTask(jctx.JoinTask)

	return j
}
