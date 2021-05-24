package cliparser

import (
	"fmt"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/iancoleman/strcase"
	"github.com/jessevdk/go-flags"
	"reflect"
)

//const CMD_NAME_DB_MIGRATE = "db_migrate"

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
