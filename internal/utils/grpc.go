package utils

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/metadata"
)

func AddMetaValue(md *metadata.MD, key string, value string) {
	md.Append(key, value)
}

func AddBinMetaValue(md *metadata.MD, key string, value []byte) {
	md.Append(key + "-bin", string(value))
}

func SetSessionId(md *metadata.MD, id g.SessionId) {
	AddMetaValue(md, g.META_FIELD_SESSION_ID, IntToHex(id, 16))
}

func GetSessionId(md *metadata.MD) g.SessionId {
	vals := md.Get(g.META_FIELD_SESSION_ID)
	if len(vals) == 0 {
		return 0
	}
	return g.SessionId(Hex2int(vals[0]))
}
