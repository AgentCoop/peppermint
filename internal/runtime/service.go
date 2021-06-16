package runtime

import "google.golang.org/protobuf/runtime/protoimpl"

type Service interface {
	Name() string
	Server() BaseServer
	IpcServer() BaseServer
	Policy() ServicePolicy
	Configurator() ServiceConfigurator
	RegisterEncKeyStoreFallback()
}

type ProtoReflect interface {
	WasSet(*protoimpl.ExtensionInfo) bool
}

// Options that can be specified both in service or method scope.
// Option value specified in a method scope will shadow the value specified in service scope.
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
	DefaultPort() int
	Ipc_UnixDomainSocket() string
	FindMethodByName(string) (Method, bool)
}

type Method interface {
	ProtoReflect
	Name() string
	ServicePolicy() ServicePolicy
	CallPolicy() MethodCallPolicy
}

type MethodCallPolicy interface {
	ServiceGlobalOptions
	IsStreamable() bool
	OpenNewSession() int
	RequiredRoles() []string
}
