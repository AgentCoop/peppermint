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

func (o *optValue) setValue(v interface{}) {
	o.isSet = true
	reflect.ValueOf(o.val).Elem().Set(reflect.ValueOf(v))
}

func (o *optValue) WasSet() bool {
	return o.isSet
}

func (o *optValue) Value() interface{} {
	return o.val
}

type (
	extOptionMap          map[*protoimpl.ExtensionInfo]*optValue
	SvcLevelOptionsMap    extOptionMap
	MethodLevelOptionsMap map[string]extOptionMap
)

func (o extOptionMap) OverrideIfNotSet(v interface{}, targetExt *protoimpl.ExtensionInfo) {
	for ext, val := range o {
		if ext == targetExt && !val.isSet {
			val.setValue(v)
		}
	}
}

func (m SvcLevelOptionsMap) AddItem(ext *protoimpl.ExtensionInfo, v interface{}) {
	m[ext] = &optValue{false, v}
}

func (m MethodLevelOptionsMap) AddItem(methodName string, ext *protoimpl.ExtensionInfo, v interface{}) {
	m[methodName][ext] = &optValue{false, v}
}

func (d ServiceDescriptor) FetchServiceCustomOptions(svcOpts SvcLevelOptionsMap) {
	m := d.Options()
	for ext, _ := range svcOpts {
		opt := svcOpts[ext]
		if !proto.HasExtension(m.(proto.Message), ext) {
			continue
		}
		v := proto.GetExtension(m.(proto.Message), ext)
		opt.setValue(v)
	}
}

func (d ServiceDescriptor) FetchMethodsCustomOptions(methodOpts MethodLevelOptionsMap) {
	ml := d.Methods().Len()
	for i := 0; i < ml; i++ {
		m := d.Methods().Get(i)
		fullName := string(m.FullName())
		parts := strings.Split(fullName, ".")
		shortName := parts[len(parts)-1]
		desc := m.Options()
		if _, has := methodOpts[shortName]; !has {
			continue
		}
		for ext, opt := range methodOpts[shortName] {
			m := desc.(proto.Message)
			if !proto.HasExtension(m, ext) {
				continue
			}
			v := proto.GetExtension(m, ext)
			opt.setValue(v)
		}
	}
}
