package cli_client

type CommonOptions struct {
	Host string `long:"host" short:"h"`
	Port uint16 `long:"port" short:"p"`
}
