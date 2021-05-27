package cmd

const (
	CMD_NAME_VERSION    = "version"
	CMD_NAME_DB_MIGRATE = "db_migrate"
	CMD_NAME_DB_CREATE  = "db_create"
	CMD_NAME_RUN        = "run"
	CMD_NAME_JOIN       = "join"
	CMD_NAME_WEB_PROXY  = "proxy-cfg"
)

type Version struct {
	Verbose []bool `short:"v" long:"verbose" description:""`
}

type Join struct {
	Tags []string `long:"tag"`
	Hub  string   `long:"hub" required:"true"`
}

type DbCreate struct {
	Force bool `long:"force" description:"Force to drop existing database before creation"`
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
		DbMigrate `command:"db_migrate"`
		DbCreate  `command:"db_create"`
		Run       `command:"run"`
		Join      `command:"join"`
		Version   `command:"version"`
	}{}
)
