package runtime

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/service"
)

// Orchestrates communication between the gRPC and Service layer
type GrpcServiceCommunicator interface {
	GrpcTx(int, interface{})
	GrpcTxStreamable(int, interface{})
	GrpcRx(int) interface{}
	ServiceTx(int, interface{})
	ServiceRx(int) interface{}
	Job() job.Job
	SessionId() i.SessionId
}

type CliParser interface {
	Data() interface{}
	Run() error
	CurrentCmd() (string, bool)
	OptionValue(string) (interface{}, bool)
	GetCmdOptions(cmdName string) (interface{}, error)
}

type Configurator interface {
	Fetch() // fetch configuration data from DB
	MergeCliOptions(CliParser)
}

type ServiceInfo struct {
	Name string
	Cfg Configurator
	Initializer func() service.Service
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


