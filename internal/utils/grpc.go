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

