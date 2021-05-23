package cliparser

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/iancoleman/strcase"
	"github.com/jessevdk/go-flags"
	"reflect"
)

const CREATEDB_CMD_NAME = "create_db"

type CmdHook func(data interface{})
//type CreateDbHook func(opts CreateDbOptions)

type parser struct {
	data interface{}
	handle *flags.Parser
}

func NewParser(data interface{}) *parser {
	p := new(parser)
	p.data = data
	p.handle = flags.NewParser(data, flags.IgnoreUnknown)
	//p.handle.SubcommandsOptional = true
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

func (p *parser) Data() interface{} {
	return p.data
}

func (p *parser) getCmdOptions(cmdName string) interface{} {
	r := reflect.ValueOf(p.data)
	fieldName := strcase.ToCamel(cmdName)
	cmdOpts := reflect.Indirect(r).FieldByName(fieldName)
	if ! cmdOpts.CanInterface() {
		panic("")
	}
	return cmdOpts.Interface()
}

func (p *parser) invokeCmdHooks(cmdName string) {
	opts := p.getCmdOptions(cmdName)
	for _, hook := range runtime.GlobalRegistry().LookupParserCmdHook(cmdName) {
		hook(opts)
	}
}
