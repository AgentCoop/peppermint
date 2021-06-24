package pkg

type CliParser interface {
	Data() interface{}
	Run() error
	CurrentCmd() (string, bool)
	OptionValue(string) (interface{}, bool)
	GetCmdOptions(cmdName string) (interface{}, error)
}

