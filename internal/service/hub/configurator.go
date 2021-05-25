package hub

import (
	model "github.com/AgentCoop/peppermint/internal/model/hub"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"net"
	"strconv"
)

type cfg struct {
	port int
	address string
	secret string
}

func NewConfigurator() *cfg {
	cfg := &cfg{}
	return cfg
}

func (c *cfg) Fetch() {
	db := runtime.GlobalRegistry().Db()
	rec := &model.HubConfig{}
	db.Handle().FirstOrCreate(rec)
	c.port = rec.Port
	c.address = rec.Address
	c.secret = rec.Secret
}

func (c *cfg) MergeCliOptions(parser runtime.CliParser) {
	// val, isset := parser.OptionValue("")
}

func (w *cfg) Address() net.Addr {
	addr, err := net.ResolveTCPAddr("tcp", w.address + ":" + strconv.Itoa(w.port))
	if err != nil { panic(err) }
	return addr
}

func (c *cfg) Secret() string {
	return c.secret
}