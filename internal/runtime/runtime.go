package runtime

import job "github.com/AgentCoop/go-work"

type CliParser interface {
	Data() interface{}
	Run() error
	CurrentCmd() (string, bool)
	GetCmdOptions(cmdName string) (interface{}, error)
}

type runtime struct {
	parser CliParser
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
	CliParser() CliParser
	InitTask(j job.Job) (job.Init, job.Run, job.Finalize)
}

func (r *runtime) CliParser() CliParser {
	return r.parser
}


