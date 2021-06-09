package server

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"github.com/AgentCoop/peppermint/internal/plugin/hub/model"
	"google.golang.org/grpc"
	"strings"
)

func isSecure(callDesc g.ServerCallDesc) (bool, []byte) {
	switch {
	// Skip insecure JoinHello call
	case strings.HasSuffix(callDesc.Method(), "JoinHello"):
		return false, nil
	// Regardless of node settings, the join procedure must use a secure channel.
	case strings.HasSuffix(callDesc.Method(), "Join"):
		node := model.FetchById(callDesc.Meta().NodeId())
		return true, node.EncKey
	default:
		node := model.FetchById(callDesc.Meta().NodeId())
		if node.E2E_EncEnabled > 0 {
			return true, node.EncKey
		} else {
			return false, nil
		}
	}
}

func decryptMessage(callDesc g.ServerCallDesc, req interface{}) {
	ok, key := isSecure(callDesc)
	if !ok { return }
	//unpacker := codec.NewUnpacker(req, codec.Serialized)
	//unpacker.Unpack(key, req)
	a := 1
	_ = a
	_ = key
}

func encryptMessage(callDesc g.ServerCallDesc, resp interface{}) {
	ok, key := isSecure(callDesc)
	if !ok { return }
	resp = codec.NewPacker(resp, codec.Serialized, key)
}

func EncLayerUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		callDesc := ctx.(g.ServerCallDesc)
		decryptMessage(callDesc, req)
		resp, err := handler(callDesc, req)
		encryptMessage(callDesc, resp)
		return resp, err
	}
}

//func EncLayerStreamInterceptor(svcName string) grpc.StreamServerInterceptor {
//	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
//		callDesc := ctx.(g.ServerCallDesc)
//
//		extended := stream.NewServerStream(ss, callDesc)
//		err := handler(srv, extended)
//		return err
//	}
//}
//
