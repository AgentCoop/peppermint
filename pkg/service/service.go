package service

import (
	"github.com/AgentCoop/peppermint/pkg/grpc"
	"github.com/AgentCoop/peppermint/pkg/node"
	"google.golang.org/protobuf/runtime/protoimpl"
	"net"
)

type Service interface {
	Name() string
	Server() grpc.BaseServer
	IpcServer() grpc.BaseServer
	Policy() ServicePolicy
	Configurator() ServiceConfigurator
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
	OpenNewSession() int
	CloseSession() bool
	RequiredRoles() []string
	Timeout() uint32
}
