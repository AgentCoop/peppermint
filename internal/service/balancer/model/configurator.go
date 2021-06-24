package model

import (
	"github.com/AgentCoop/peppermint/internal/runtime"
	"github.com/AgentCoop/peppermint/pkg"
	"github.com/AgentCoop/peppermint/pkg/service"
	"github.com/AgentCoop/peppermint/pkg/service/balancer"
	"net"
	"strconv"
)

type HubConfigurator interface {
	service.ServiceConfigurator
	Secret() string
}

type prefAlgoMap map[string]balancer.Algo

type cfg struct {
	port          int
	address       string
	defaultAlgo   balancer.Algo
	preferredAlgo prefAlgoMap
}

func NewConfigurator() *cfg {
	cfg := &cfg{}
	return cfg
}

func (c *cfg) Fetch() error {
	db := runtime.GlobalRegistry().Db().Handle()
	rec := BalancerConfig{}
	err := db.FirstOrCreate(&rec).Error
	if err != nil {
		return err
	}
	c.port = rec.Port
	c.address = rec.Address
	c.defaultAlgo = balancer.Algo(rec.DefaultAlgo)
	c.preferredAlgo = make(prefAlgoMap)
	for _, v := range rec.PreferredAlgo {
		c.preferredAlgo[v.ServiceName] = balancer.Algo(v.Algo)
	}
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
	addr, err := net.ResolveTCPAddr("tcp", c.address+":"+strconv.Itoa(c.port))
	if err != nil {
		panic(err)
	}
	return addr
}

func (c *cfg) DefaultAlgo() balancer.Algo {
	return c.DefaultAlgo()
}

func (c *cfg) PreferredAlgoByServiceName(svcName string) (balancer.Algo, bool) {
	v, ok := c.preferredAlgo[svcName]
	return v, ok
}
