package runtime

import (
	"context"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	//"github.com/AgentCoop/peppermint/internal/service"
	"net"
)

type NodeStatus int
const (
	Available NodeStatus = iota + 1
)

type Node interface {
	Id() i.NodeId
	ServiceEndpointByName(string) ServiceEndpoint
}

type NodePool interface {
	Add(Node)
	Remove(i.NodeId)
	FindById(i.NodeId) Node
	FilterByStatus(NodeStatus) NodePool
	Len() int
}

type ServiceLocator interface {
	FindByMethodName(string) NodePool
	ServiceNameByMethod(string) string
}

type ServiceEndpoint interface {
	Address() net.Addr
	EncKey() []byte
}

type Stream interface {
	Context() context.Context
	SendMsg(m interface{}) error
	RecvMsg(m interface{}) error
}

type StreamInfo interface {
	EncKey() []byte
	FullMethod() string
	MessagesReceived() int
}

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

type ServiceConfigurator interface {
	Configurator
	Address() net.Addr
}

type Service interface {
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
}

type ServiceInfo struct {
	Name string
	Cfg Configurator
	Initializer func() Service
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


