package stream

import (
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"google.golang.org/grpc"
	"io"
)

type clientStream struct {
	cs         grpc.ClientStream
	callDesc   g.ClientCallDesc
	recvx      int
	sentx      int
}

func (stream *clientStream) Close() {
	stream.cs.SendMsg(io.EOF)
}

func (stream *clientStream) SendMsg(msg interface{}) error {
	var err error
	switch v := msg.(type) {
	case error:
		if v == io.EOF {
			stream.cs.CloseSend()
			return nil
		} else {
			return v
		}
	default:
		if stream.sentx == 0 {
			stream.callDesc.Meta().SendHeader(nil)
			//grpc.SendHeader(stream.Context(), *stream.header)
		}
		err = stream.cs.SendMsg(msg)
		stream.sentx++
		return err
	}
}

func (stream *clientStream) RecvMsg(msg interface{}) error {
	var err error
	err = stream.cs.RecvMsg(msg)
	switch err {
	case nil:
		stream.recvx++
		return nil
	case io.EOF:
		return nil
	default:
		return err
	}
}

func (stream *clientStream) ReceivedCount() int {
	return stream.recvx
}

func (stream *clientStream) SentCount() int {
	return stream.sentx
}
