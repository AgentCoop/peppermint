package runtime

type HookType int
type Hook func()

const (
	HOOK_CLI_ARG HookType = iota
	HOOK_BEFORE_SERVICE_START
)

func AddHook(typ HookType, hook Hook) {

}
