package model

import (
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
	hubDb   hubDb
	port    int
	address string
	secret  string
}

func (c *cfg) Fetch(nodeId uint) error {
	db := c.hubDb.Handle()
	rec := &HubConfig{}
	err := db.FirstOrCreate(rec, nodeId).Error
	if err != nil {
		return err
	}
	c.port = rec.Port
	c.address = rec.Address
	c.secret = rec.Secret
	return nil
}

func (c *cfg) Refresh(nodeId uint) error {
	return c.Fetch(nodeId)
}

func (c *cfg) MergeCliOptions(parser pkg.CliParser) {
	val, isset := parser.OptionValue("hub-port")
	if isset {
		c.port = val.(int)
	}
}

func (c *cfg) Address() net.Addr {
	addr, err := net.ResolveTCPAddr("tcp", c.address+":"+strconv.Itoa(c.port))
	if err != nil {
		panic(err)
	}
	return addr
}

func (c *cfg) Secret() string {
	return c.secret
}
