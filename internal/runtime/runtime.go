package runtime

import job "github.com/AgentCoop/go-work"

type CliParser interface {
	Data() interface{}
	Run() error
}

type runtime struct {
	CliParser
	dbFilename string
}

func NewRuntime(parser CliParser, dbFilename string) *runtime {
	r := &runtime{
		parser,
		dbFilename,
	}
	GlobalRegistry().SetRuntime(r)
	return r
}

type Runtime interface {
	InitTask(j job.Job) (job.Init, job.Run, job.Finalize)
	//OsSignalTask(j job.Job) (job.Init, job.Run, job.Finalize)
	NodeTask(j job.Job) (job.Init, job.Run, job.Finalize)
}


