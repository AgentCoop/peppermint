package utils

import (
	i "github.com/AgentCoop/peppermint/internal"
	grpc "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AddMetaValue(md *metadata.MD, key string, value string) {
	md.Append(key, value)
}

func AddBinMetaValue(md *metadata.MD, key string, value []byte) {
	md.Append(key + "-bin", string(value))
}

func Grpc_SessionId(md *metadata.MD, id i.SessionId) {
	if id == 0 { return }
	AddMetaValue(md, grpc.META_FIELD_SESSION_ID, Conv_IntToHex(id, 16))
}

func Grpc_ExtractSessionId(md *metadata.MD) i.SessionId {
	vals := md.Get(grpc.META_FIELD_SESSION_ID)
	if len(vals) == 0 { return 0 }
	return i.SessionId(Conv_HexToInt(vals[0]))
}

func GrpcErrorWrapper(err interface{}) interface{} {
	var text string
	switch v := err.(type) {
	case error:
		text = err.(error).Error()
	case string:
		text = v
	case status.Status:
		return v
	default:
		text = "Unknown system error"
	}
	return status.Error(codes.Internal, text)
}

