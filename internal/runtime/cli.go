package runtime

import (
	"github.com/jessevdk/go-flags"
)

type CliParser interface {
	Run()
	AddCommand(name string, options interface{}, short string, long string)
}

type cliParser struct {
	parser *flags.Parser
}

func NewCliParser(options interface{}) *cliParser {
	p := &cliParser{}
	p.parser = flags.NewParser(options, flags.PassDoubleDash | flags.PrintErrors | flags.IgnoreUnknown)
	return p
}

func (c *cliParser) AddCommand(name string, options interface{}, short string, long string) {
	cmd := c.parser.Find(name)
	if cmd != nil {
		return
	}
	c.parser.AddCommand(name, short, long, options)
}

func (c *cliParser) Run() {
	c.parser.Parse()
}
