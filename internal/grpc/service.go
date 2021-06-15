package grpc

type Service interface {
	Name() string
	Server() BaseServer
	IpcServer() BaseServer
	Policy() ServicePolicy
	RegisterEncKeyStoreFallback()
}

type ServicePolicy interface {
	EnforceEncryption() bool
	DefaultPort() int
	Ipc_UnixDomainSocket() string
	FindMethodByName(string) (Method, bool)
}

type Method interface {
	Name() string
	CallPolicy() MethodCallPolicy
}

type MethodCallPolicy interface {
	IsSecure() bool
	IsStreamable() bool
	OpenNewSession() int
	RequiredRoles() []string
}
