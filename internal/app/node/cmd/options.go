package cmd

const (
	CMD_NAME_BOOTSTRAP  = "bootstrap"
	CMD_NAME_VERSION    = "version"
	CMD_NAME_DB_MIGRATE = "db_migrate"
	CMD_NAME_RUN        = "run"
	CMD_NAME_JOIN       = "join"
	CMD_NAME_WEB_PROXY  = "proxy-cfg"
)

type netOptions struct {
	Port   uint16 `long:"port" short:"p"`
	Target string `long:"target" required:"true"`
}

type Bootstrap struct {
	Tags            []string `long:"tag"`
	IdFromInterface string   `long:"id-from-nic"`
	CreateDb        bool     `long:"create-db"`
	Force           bool     `long:"force"`
}

type Version struct {
	Verbose []bool `short:"v" long:"verbose" description:""`
}

type Join struct {
	Tags []string `long:"tag"`
	Hub  string   `long:"hub" required:"true"`
}

type DbMigrate struct {
	Drop bool `long:"drop" description:"Drop database before migration"`
}

type Run struct {
	HubPort      int `long:"hub-port"`
	WebProxyPort int `long:"wp-port"`
}

var (
	Options = struct {
		AppDir    string `long:"appdir" short:"d"`
		NodeId    uint   `long:"node-id" short:"n"`
		Bootstrap `command:"bootstrap"`
		DbMigrate `command:"db_migrate"`
		Run       `command:"run"`
		Join      `command:"join"`
		Version   `command:"version"`
	}{}
)
