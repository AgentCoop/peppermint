package service

import (
	job "github.com/AgentCoop/go-work"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/grpc"
	"github.com/AgentCoop/peppermint/pkg/node"
	"google.golang.org/protobuf/runtime/protoimpl"
	"net"
)

type Service interface {
	Name() string
	ShortName() string
	Db() pkg.Db
	OpenDb() error
	Configurator() ServiceConfigurator
	WithConfigurator(ServiceConfigurator)
	ReloadConfig(uint) error
	FetchConfig(uint) (ServiceConfigurator, error)
	Server() grpc.BaseServer
	WithServer(grpc.BaseServer)
	IpcServer() grpc.BaseServer
	WithIpcServer(server grpc.BaseServer)
	Policy() ServicePolicy
	WithPolicy(ServicePolicy)
	InitTask(j job.Job) (job.Init, job.Run, job.Finalize)
}

type ServiceConfigurator interface {
	node.Configurator
	Address() net.Addr
}

type ProtoReflect interface {
	WasSet(*protoimpl.ExtensionInfo) bool
}

// Options that can be specified both in service or method scope.
// Option value specified in a method scope will override the value specified in the service scope.
type ServiceGlobalOptions interface {
	EnforceEncryption() bool
}

type ServiceOptionValue interface {
	WasSet() bool
	Value() interface{}
}

type ServicePolicy interface {
	ProtoReflect
	ServiceGlobalOptions
	DefaultPort() uint16
	Ipc_UnixDomainSocket() string
	FindMethodByName(string) (Method, bool)
}

type Method interface {
	ProtoReflect
	Name() string
	FullName() string
	ServicePolicy() ServicePolicy
	CallPolicy() MethodCallPolicy
}

type MethodCallPolicy interface {
	ServiceGlobalOptions
	IsStreamable() bool
	SessionSticky() bool
	NewSession() int
	CloseSession() bool
	RequiredRoles() []string
	Timeout() uint32
}
