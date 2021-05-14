package cmd

import (
	"github.com/jessevdk/go-flags"
)

type CliParser struct {
	parser *flags.Parser
}

func NewCliParser(options interface{}) *CliParser {
	p := &CliParser{}
	p.parser = flags.NewParser(options, flags.PassDoubleDash | flags.PrintErrors | flags.IgnoreUnknown)
	return p
}

func (c *CliParser) AddCommand(name string, options interface{}) {
	cmd := c.parser.Find(name)
	if cmd != nil {
		return
	}
	c.parser.AddCommand(name, "", "", options)
}
