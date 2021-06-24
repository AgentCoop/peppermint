package model

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/service"
	"net"
	"strconv"
)

type HubConfigurator interface {
	service.ServiceConfigurator
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
	db := runtime.GlobalRegistry().Db().Handle()
	rec := HubConfig{}
	err := db.FirstOrCreate(&rec).Error
	if err != nil { return err }
	//errors.Is(err, gorm.ErrRecordNotFound)
	c.port = rec.Port
	c.address = rec.Address
	c.secret = rec.Secret
	return nil
}

func (c *cfg) Refresh() error {
	return c.Fetch()
}

func (c *cfg) MergeCliOptions(parser pkg.CliParser) {
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