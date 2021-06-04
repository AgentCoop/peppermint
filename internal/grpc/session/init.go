package session

//import "github.com/AgentCoop/peppermint/internal/runtime"

func init() {
	sMap = make(sessionMap, 0)
	//runtime.GlobalRegistry().SetGrpcSession(sMap)
}
