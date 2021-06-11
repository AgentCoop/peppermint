package grpc

type Service interface {
	Name() string
	Server() BaseServer
	Policy() ServicePolicy
	Methods() Method
	RegisterEncKeyStoreFallback()
}

type ServicePolicy interface {
	EnforceEncryption() bool
	DefaultPort() int
	Ipc_UnixDomainSocket() string
}

type Method interface {
	Name() string
	FullName() string
	CallPolicy() MethodCallPolicy
}

type MethodCallPolicy interface {
	IsSecure() bool
	IsStreamable() bool
	OpenNewSession() int
	RequiredRoles() []string
}
