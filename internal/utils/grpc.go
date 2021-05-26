package utils

import (
	job "github.com/AgentCoop/go-work"
	i "github.com/AgentCoop/peppermint/internal"
	grpc "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/runtime"
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

func GetSessDescriptorById(id i.SessionId) (runtime.SessionDesc, error) {
	sess := runtime.GlobalRegistry().GrpcSession()
	desc, ok := sess.Lookup(id)
	if !ok  {
		return nil, status.Error(codes.DeadlineExceeded, "gRPC session has been expired or session ID is invalid")
	} else {
		return desc, nil
	}
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
	return status.New(codes.Internal, text)
}

func DefaultGrpcJob(value interface{}) job.Job {
	job := job.NewJob(value)
	job.WithErrorWrapper(GrpcErrorWrapper)
	return job
}