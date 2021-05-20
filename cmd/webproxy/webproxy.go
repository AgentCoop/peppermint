package main

import (
	"fmt"
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/internal/grpc/webproxy"
)

func main() {
	j := job.NewJob(nil)

	proxy := webproxy.NewServer("localhost:9900")

	j.AddTask(proxy.StartTask)

	<-j.Run()
	_, err := j.GetInterruptedBy()
	if err != nil {
		fmt.Printf("err %v\n", err)
	}
}
