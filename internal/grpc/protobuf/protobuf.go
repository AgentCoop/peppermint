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

type optValue struct {
	isSet bool
	val   interface{}
}

type extOptionMap map[*protoimpl.ExtensionInfo]optValue
type svcLevelOptionsMap extOptionMap
type methodLevelOptionsMap map[string]extOptionMap

func (m svcLevelOptionsMap) AddItem(ext *protoimpl.ExtensionInfo, v interface{}) {
	m[ext] = optValue{false, v}
}

func (m methodLevelOptionsMap) AddItem(methodName string, ext *protoimpl.ExtensionInfo, v interface{}) {
	if _, has := m[methodName]; !has {
		m[methodName] = make(extOptionMap, 0)
	}
	m[methodName][ext] = optValue{false, v}
}

func (m methodLevelOptionsMap) OverrideVal(src optValue, target *protoimpl.ExtensionInfo) {
	for _, opts := range m {
		for ext, v := range opts {
			if target == ext && !v.isSet {
				v.val = src.val
			}
		}
	}
}

func (d ServiceDescriptor) FetchServiceCustomOptions(svcOpts svcLevelOptionsMap, methodOpts methodLevelOptionsMap) {
	m := d.Options()
	for ext, _ := range svcOpts {
		opt := svcOpts[ext]
		if !proto.HasExtension(m.(proto.Message), ext) {
			opt.isSet = false
			continue
		}
		v := proto.GetExtension(m.(proto.Message), ext)
		reflect.ValueOf(opt.val).Elem().Set(reflect.ValueOf(v))
	}
	ml := d.Methods().Len()
	for i := 0; i < ml; i++ {
		m := d.Methods().Get(i)
		fullName := string(m.FullName())
		parts := strings.Split(fullName, ".")
		shortName := parts[len(parts)-1]
		msg := m.Options()
		if _, has := methodOpts[shortName]; !has {
			continue
		}
		for ext, optVal := range methodOpts[shortName] {
			m := msg.(proto.Message)
			if !proto.HasExtension(m, ext) {
				optVal.isSet = false
				continue
			}
			optVal.isSet = true
			v := proto.GetExtension(m, ext)
			reflect.ValueOf(optVal.val).Elem().Set(reflect.ValueOf(v))
		}
	}
}
