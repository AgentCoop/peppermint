package protobuf

import (
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"reflect"
	"strings"
)

type ServiceDescriptor struct {
	protoreflect.ServiceDescriptor
}

type extValueMap = map[*protoimpl.ExtensionInfo]interface{}
type ServiceLevelOptions extValueMap
type MethodLevelOptions map[string]extValueMap

func (m MethodLevelOptions) AddItem(methodName string, k *protoimpl.ExtensionInfo, v interface{}) {
	if _, ok := m[methodName]; !ok {
		m[methodName] = make(extValueMap, 0)
	}
	m[methodName][k] = v
}

func (d ServiceDescriptor) FetchServiceCustomOptions(svcOpts ServiceLevelOptions, methodOpts MethodLevelOptions) {
	m := d.Options()
	for ext, _ := range svcOpts {
		v := proto.GetExtension(m.(proto.Message), ext)
		recv := svcOpts[ext]
		reflect.ValueOf(recv).Elem().Set(reflect.ValueOf(v))
	}
	var (
		mOptsMap extValueMap
		has      bool
	)
	ml := d.Methods().Len()
	for i := 0; i < ml; i++ {
		m := d.Methods().Get(i)
		fullName := string(m.FullName())
		parts := strings.Split(fullName, ".")
		shortName := parts[len(parts)-1]
		msg := m.Options()
		if mOptsMap, has = methodOpts[shortName]; !has { continue }
		for ext, _ := range mOptsMap {
			v := proto.GetExtension(msg.(proto.Message), ext)
			recv := mOptsMap[ext]
			reflect.ValueOf(recv).Elem().Set(reflect.ValueOf(v))
		}
	}
}