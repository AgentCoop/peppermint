package webproxy

import (
	model "github.com/AgentCoop/peppermint/internal/service/webproxy/model"
	"github.com/AgentCoop/peppermint/internal/runtime"
	"io/ioutil"
	"net"
	"strconv"
)

type cfg struct {
	model model.WebProxyConfig
	port int
	address string
}

func NewConfigurator() *cfg {
	return &cfg{}
}

func (w *cfg) Fetch() error {
	db := runtime.GlobalRegistry().Db()
	rec := &model.WebProxyConfig{}
	db.Handle().FirstOrCreate(rec)
	w.port = rec.Port
	w.address = "peppermint.io" //rec.Address
	return nil
}

func (w *cfg) MergeCliOptions(parser runtime.CliParser) {
	val, isset := parser.OptionValue("wp-port")
	if isset {
		w.port = val.(int)
	}
}

func (w *cfg) Address() net.Addr {
	addr, err := net.ResolveTCPAddr("tcp", w.address + ":" + strconv.Itoa(w.port))
	if err != nil { panic(err) }
	return addr
}

func (w *cfg) ServerName() string {
	return "peppermint.io"
}

func (w *cfg) X509CertPEM() []byte {
	cert, _ := ioutil.ReadFile("/home/pihpah/mycerts/peppermint.io/server.crt")
	return cert
}

func (w *cfg) X509KeyPEM() []byte {
	key, _ := ioutil.ReadFile("/home/pihpah/mycerts/peppermint.io/server.key")
	return key
}
