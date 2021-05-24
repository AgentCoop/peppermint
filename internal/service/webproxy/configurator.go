package webproxy

import (
	model "github.com/AgentCoop/peppermint/internal/model/webproxy"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"net"
)

type cfg struct {
	model model.WebProxyConfig
	port int
	Foo string
	address string
}

func NewConfigurator() *cfg {
	return &cfg{}
}

func (w *cfg) Fetch() {
	db := runtime.GlobalRegistry().Db()
	rec := &model.WebProxyConfig{}
	db.Handle().FirstOrCreate(rec)
	w.port = rec.Port
	w.address = rec.Address
}

func (w *cfg) MergeCliOptions(parser runtime.CliParser) {
	_, isset := parser.OptionValue("wp-port")
	if isset {
		w.port = 12000 //val.(int)
	}
}

func (w *cfg) Address() net.Addr {
	return &net.TCPAddr{
		IP:   []byte(w.address),
		Port: w.port,
		Zone: "",
	}
}

func (w *cfg) ServerName() string {
	return ""
}

func (w *cfg) X509CertPEM() []byte {
	return nil
}

func (w *cfg) X509KeyPEM() []byte {
	return nil
}
