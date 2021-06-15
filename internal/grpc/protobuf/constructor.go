package protobuf

import (
	"fmt"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"reflect"
)

func NewServiceDescriptor(svcFullName string) ServiceDescriptor {
	fullName := protoreflect.FullName(svcFullName)
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(fullName)
	if err != nil { panic(err) }
	sd, ok := desc.(protoreflect.ServiceDescriptor)
	if !ok {
		got, wanted := reflect.TypeOf(sd).String(), "protoreflect.ServiceDescriptor"
		panic(fmt.Errorf("protobuf: got %s, wanted %s", got, wanted))
	}
	return ServiceDescriptor{sd}
}

func NewMethodLevelOptions(methods []string) methodLevelOptionsMap {
	opts := make(methodLevelOptionsMap, 0)
	for _, name := range methods {
		opts[name] = make(extOptionMap, 0)
	}
	return opts
}

func NewSvcLevelOptions() svcLevelOptionsMap {
	m := make(svcLevelOptionsMap)
	return m
}