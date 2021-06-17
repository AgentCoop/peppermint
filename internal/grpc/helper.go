package grpc

import (
	i "github.com/AgentCoop/peppermint/internal"
	"github.com/AgentCoop/peppermint/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func ErrorWrapper(err interface{}) interface{} {
	var text string
	switch v := err.(type) {
	case error:
		_, ok := status.FromError(v)
		if ok { // return as is
			return v
		}
		text = v.Error()
	case string:
		text = v
	default:
		text = "Unknown system error"
	}
	return status.Error(codes.Internal, text)
}

func appendBin(md *metadata.MD, key string, value []byte) {
	md.Append(key + "-bin", string(value))
}

func SetSessionId(md *metadata.MD, id i.SessionId) {
	if id == 0 { return }
	md.Append(META_FIELD_SESSION_ID, utils.Conv_IntToHex(id, 16))
}

func SetNodeId(md *metadata.MD, id i.NodeId) {
	if id == 0 { return }
	md.Append(META_FIELD_NODE_ID, utils.Conv_IntToHex(id, 16))
}

func ExtractSessionId(md *metadata.MD) i.SessionId {
	vals := md.Get(META_FIELD_SESSION_ID)
	if len(vals) == 0 { return 0 }
	return i.SessionId(utils.Conv_HexToInt(vals[0]))
}

func ExtractNodeId(md *metadata.MD) i.NodeId {
	vals := md.Get(META_FIELD_NODE_ID)
	if len(vals) == 0 { return 0 }
	return i.NodeId(utils.Conv_HexToInt(vals[0]))
}
