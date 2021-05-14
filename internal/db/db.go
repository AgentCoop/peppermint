package db

import job "github.com/AgentCoop/go-work"

type SqlTransaction interface {
	Start()
	Rollback()
	Commit()
	AddQuery()
}

type SqlQuery interface {

}

type TaskHook func()

type unitOfWorkContext struct {
	sqlTrans SqlTransaction
	hook TaskHook
}

func (c *unitOfWorkContext) SqlQueryExecTask(j job.Job) (job.Init, job.Run, job.Finalize) {
	init := func(task job.Task) {
		c.sqlTrans.Start()
	}
	run := func(task job.Task) {
		c.hook()
		task.Done()
	}
	fin := func(task job.Task) {
		switch j.GetState() {
		case job.Cancelled:
			c.sqlTrans.Rollback()
		default:
			c.sqlTrans.Commit()
		}
	}
	return init, run, fin
}

