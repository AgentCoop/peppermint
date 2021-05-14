package runtime

type Runtime interface {
	Init()
}

type runtime struct {
	cliParser CliParser
}

var (
	rt Runtime
)

func NewRuntime(parser CliParser) *runtime {
	r := &runtime{}
	r.cliParser = parser
	return r
}

func (r *runtime) Init() {
	r.cliParser.Run()
}

func GetRuntime() Runtime {
	return rt
}
