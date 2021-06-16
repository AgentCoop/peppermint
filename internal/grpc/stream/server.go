package stream

import (
	"context"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
)

type serverStream struct {
	ss         grpc.ServerStream
	callDesc   g.ServerDescriptor
	recvx      int
	sentx      int
	fullMethod string
}

func (stream *serverStream) Context() context.Context {
	return stream.ss.Context()
}

func (stream *serverStream) Close() {
	stream.SendMsg(io.EOF)
}

func (stream *serverStream) SendMsg(msg interface{}) error {
	var err error
	sec := stream.callDesc.SecPolicy()
	err = encLayer(msg, sec.IsSecure(), sec.EncKey())
	if err != nil { return nil }

	if msg == io.EOF && stream.callDesc.Meta().Trailer() != nil {
		stream.callDesc.Meta().SetTrailer(*stream.callDesc.Meta().Trailer())
		return nil
	}
	if stream.sentx == 0 {
		stream.callDesc.Meta().SendHeader(nil)
	}
	stream.sentx++
	err = stream.ss.SendMsg(msg)
	return err
}

func (stream *serverStream) RecvMsg(msg interface{}) error {
	var err error
	sec := stream.callDesc.SecPolicy()
	err = encLayer(msg, sec.IsSecure(), sec.EncKey())
	if err != nil { return nil }

	err = stream.ss.RecvMsg(msg)
	if err == io.EOF {
		return nil
	}
	stream.recvx++
	return err
}

func (stream *serverStream) FullMethod() string {
	return stream.fullMethod
}

func (stream *serverStream) ReceivedCount() int {
	return stream.recvx
}

func (stream *serverStream) SentCount() int {
	return stream.sentx
}

//
// ServerStream compatibility layer
//
func (s *serverStream) SetHeader(md metadata.MD) error {
	panic("implement me")
}

func (s *serverStream) SendHeader(md metadata.MD) error {
	panic("implement me")
}

func (s *serverStream) SetTrailer(md metadata.MD) {
	panic("implement me")
}
