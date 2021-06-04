package stream

import (
	"context"
	"errors"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
)

var (
	ErrEmptyEncryptionKey       = errors.New("stream: encryption key is empty")
	ErrFailedToRetrieveMetadata = errors.New("stream: failed to retrieve metadata")
)

type streamType int

const (
	ClientStream streamType = iota
	ServerStream
)

type stream struct {
	g.Stream
	header, trailer *metadata.MD
	typ             streamType
	isSecure        bool
	encKey          []byte
	recvx           int
	sentx           int
	fullMethod      string
}

func (s *stream) encLayer(msg interface{}) error {
	_, ok := msg.(codec.Packet)
	if ok { return nil }
	if s.isSecure {
		if len(s.encKey) == 0 {
			return ErrEmptyEncryptionKey
		}
		msg = codec.NewPacket(msg, s.encKey)
	}
	return nil
}

func (s *stream) Context() context.Context {
	return s.Context()
}

func (s *stream) Close() {
	s.Send(io.EOF)
}

func (s *stream) Send(msg interface{}) error {
	err := s.encLayer(msg)
	if err != nil { return nil }
	switch s.typ {
	case ServerStream:
		return s.srvSend(msg)
	case ClientStream:
		return s.clientSend(msg)
	}
	return nil
}

func (s *stream) Recv(msg interface{}) error {
	err := s.encLayer(msg)
	if err != nil { return nil }
	switch s.typ {
	case ServerStream:
		return s.srvRecv(msg)
	case ClientStream:
		return s.clientRecv(msg)
	}
	return nil
}

func (s *stream) srvSend(msg interface{}) error {
	ss := s.Stream.(grpc.ServerStream)
	var err error
	if msg == io.EOF && s.trailer != nil {
		ss.SetTrailer(*s.trailer)
		return nil
	}
	if s.sentx == 0 {
		ss.SendHeader(*s.header)
	}
	s.sentx++
	err = ss.SendMsg(msg)
	return err
}

func (s *stream) srvRecv(msg interface{}) error {
	ss := s.Stream.(grpc.ServerStream)
	var err error
	err = ss.RecvMsg(msg)
	if err == io.EOF {
		return nil
	}
	s.recvx++
	return err
}

func (s *stream) clientRecv(msg interface{}) error {
	cs := s.Stream.(grpc.ClientStream)
	var err error
	err = cs.RecvMsg(msg)
	switch err {
	case nil:
		s.recvx++
		return nil
	case io.EOF:
		return nil
	default:
		return err
	}
}

func (s *stream) clientSend(msg interface{}) error {
	cs := s.Stream.(grpc.ClientStream)
	var err error
	switch v := msg.(type) {
	case error:
		if v == io.EOF {
			cs.CloseSend()
			return nil
		} else {
			return v
		}
	default:
		if s.sentx == 0 {
			grpc.SendHeader(s.Context(), *s.header)
		}
		err = cs.SendMsg(msg)
		s.sentx++
		return err
	}
}

func (s *stream) Header() metadata.MD {
	var md metadata.MD
	var ok bool
	switch s.typ {
	case ServerStream:
		md, ok = metadata.FromIncomingContext(s.Stream.Context())
		return md
	case ClientStream:
		md, ok = metadata.FromOutgoingContext(s.Stream.Context())
	}
	if !ok {
		panic(ErrFailedToRetrieveMetadata)
	}
	return md.Copy()
}

func (s *stream) EncKey() []byte {
	return s.encKey
}

func (s *stream) WithNewHeader(md *metadata.MD) {
	s.header = md
}

func (s *stream) WithTrailer(md *metadata.MD) {
	s.trailer = md
}

func (s *stream) FullMethod() string {
	return s.fullMethod
}

func (s *stream) MessagesReceived() int {
	return s.recvx
}

func (s *stream) MessagesSent() int {
	return s.sentx
}
