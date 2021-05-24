package cliparser

import (
	"fmt"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/iancoleman/strcase"
	"github.com/jessevdk/go-flags"
	"reflect"
)

type CmdHook func(data interface{})

type parser struct {
	data interface{}
	cfgs []runtime.Configurator
	handle *flags.Parser
}

func NewParser(data interface{}) *parser {
	p := new(parser)
	p.data = data
	p.handle = flags.NewParser(data, flags.IgnoreUnknown)
	p.cfgs = make([]runtime.Configurator, 0)
	return p
}

func (p *parser) Run() error {
	_, err := p.handle.Parse()
	if err != nil { panic(err) }

	if p.handle.Active != nil {
		p.invokeCmdHooks(p.handle.Active.Name)
	}
	return err
}

func (p *parser) OptionValue(longName string) (interface{}, bool) {
	var opt *flags.Option
	switch {
	case p.handle.Active != nil:
		opt = p.handle.Active.FindOptionByLongName(longName)
	default:
		opt = p.handle.FindOptionByLongName(longName)
	}
	if opt == nil || opt.IsSetDefault() {
		return nil, false
	} else {
		return opt.Value(), true
	}
}

func (p *parser) Data() interface{} {
	return p.data
}

func (p *parser) CurrentCmd() (string, bool) {
	switch {
	case p.handle.Active != nil:
		return p.handle.Active.Name, true
	default:
		return "", false
	}
}

func (p *parser) GetCmdOptions(cmdName string) (interface{}, error) {
	r := reflect.ValueOf(p.data)
	fieldName := strcase.ToCamel(cmdName)
	cmdOpts := reflect.Indirect(r).FieldByName(fieldName)
	if ! cmdOpts.CanInterface() {
		return nil, fmt.Errorf("cli-parser: failed to retrieve options field %s", fieldName)
	}
	return cmdOpts.Interface(), nil
}

func (p *parser) invokeCmdHooks(cmdName string) {
	opts, _ := p.GetCmdOptions(cmdName)
	for _, hook := range runtime.GlobalRegistry().LookupParserCmdHook(cmdName) {
		hook(opts)
	}
}
