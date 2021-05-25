package utils

import (
	i "github.com/AgentCoop/peppermint/internal"
	grpc "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/metadata"
)

func AddMetaValue(md *metadata.MD, key string, value string) {
	md.Append(key, value)
}

func AddBinMetaValue(md *metadata.MD, key string, value []byte) {
	md.Append(key + "-bin", string(value))
}

func SetGrpcSessionId(md *metadata.MD, id i.SessionId) {
	if id == 0 {
		return
	}
	AddMetaValue(md, grpc.META_FIELD_SESSION_ID, IntToHex(id, 16))
}

func ExtractGrpcSessionId(md *metadata.MD) i.SessionId {
	vals := md.Get(grpc.META_FIELD_SESSION_ID)
	if len(vals) == 0 {
		return 0
	}
	return i.SessionId(Hex2int(vals[0]))
}

func RandomGrpcSessionId() i.SessionId {
	return i.SessionId(RandUint64())
}
