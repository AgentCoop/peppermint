package model

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/internal/runtime/deps"
	"net"
	"strconv"
)

type HubConfigurator interface {
	deps.ServiceConfigurator
	Secret() string
}

type cfg struct {
	port int
	address string
	secret string
}

func NewConfigurator() *cfg {
	cfg := &cfg{}
	return cfg
}

func (c *cfg) Fetch() error {
	db := runtime.GlobalRegistry().Db()
	rec := &HubConfig{}
	db.Handle().FirstOrCreate(rec)
	c.port = rec.Port
	c.address = rec.Address
	c.secret = rec.Secret
	return nil
}

func (c *cfg) MergeCliOptions(parser deps.CliParser) {
	val, isset := parser.OptionValue("hub-port")
	if isset {
		c.port = val.(int)
	}
}

func (c *cfg) Address() net.Addr {
	addr, err := net.ResolveTCPAddr("tcp", c.address + ":" + strconv.Itoa(c.port))
	if err != nil { panic(err) }
	return addr
}

func (c *cfg) Secret() string {
	return c.secret
}