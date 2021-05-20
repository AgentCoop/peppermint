package join

import job "github.com/AgentCoop/go-work"

func (ctx *joinCtx) JoinTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	run := func(task job.Task) {
		//req := <- ctx.joinReqCh
		//req.secret
	}
	return nil, run, nil
}

