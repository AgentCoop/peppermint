package server

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/session"
	"github.com/AgentCoop/peppermint/internal/grpc/stream"
	"google.golang.org/grpc"
	"time"
)

func fetchSession(desc g.ServerDescriptor) (g.Session, error) {
	callPolicy := desc.Method().CallPolicy()
	switch {
	case callPolicy.NewSession() > 0:
		sess := session.NewSession(time.Duration(callPolicy.NewSession()))
		desc.Meta().SetSessionId(sess.Id())
		return sess, nil
	case callPolicy.SessionSticky():
		sId := desc.Meta().SessionId()
		sess, err := session.FindById(sId)
		return sess, err
	default:
		return nil, nil
	}
}

func SessionUnaryInterceptor(svcName string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		desc := (ctx).(g.ServerDescriptor)
		sess, err := fetchSession(desc)
		if err != nil {
			return nil, err
		}
		desc.WithSession(sess)
		resp, err := handler(ctx, req)
		return resp, err
	}
}

func SessionStreamInterceptor(svcName string) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		callDesc := handleMeta(ss.Context(), svcName, info.FullMethod)
		extended := stream.NewServerStream(ss, callDesc)
		err := handler(srv, extended)
		return err
	}
}
