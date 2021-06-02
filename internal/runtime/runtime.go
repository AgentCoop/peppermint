package runtime

import (
	"context"
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"

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

type Service interface {
	StartTask(j job.Job) (job.Init, job.Run, job.Finalize)
}

type ServiceInfo struct {
	Name string
	Cfg deps.Configurator
	Initializer func() Service
}

type runtime struct {
	nodeCfg deps.NodeConfigurator
	parser deps.CliParser
}

func NewRuntime(nodeCfg deps.NodeConfigurator, parser deps.CliParser) *runtime {
	r := &runtime{
		nodeCfg: nodeCfg,
		parser: parser,
	}
	GlobalRegistry().SetRuntime(r)
	return r
}

type Runtime interface {
	CliParser() deps.CliParser
	NodeConfigurator() deps.NodeConfigurator
}

func (r *runtime) NodeConfigurator() deps.NodeConfigurator {
	return r.nodeCfg
}

func (r *runtime) CliParser() deps.CliParser {
	return r.parser
}
