package balancer

import (
	job "github.com/AgentCoop/go-work"
	"io"
)

func (c *proxyConn) readUpstreamTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {

	}
	run := func(task job.Task) {
		var err error
		err = c.upstream.cs.RecvMsg()
		task.Assert(err)

		if err == io.EOF {
			task.Done()
		}

		if c.msgSent == 0 {
			md, err := c.upstream.cs.Header()
			task.Assert(err)
			c.upstream.ss.SendHeader(md)
		}

		err = c.upstream.ss.SendMsg()
		task.Assert(err)
		c.msgSent++

		task.Tick()
	}
	fin := func(task job.Task) {

	}
	return init, run, fin
}

