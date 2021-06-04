package stream

import (
	"errors"
	g "github.com/AgentCoop/peppermint/internal/grpc"
	"github.com/AgentCoop/peppermint/internal/grpc/codec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
)

var (
	ErrEmptyEncryptionKey = errors.New("stream: encryption key is empty")
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

func (s *stream) Close() {
	s.Send(io.EOF)
}

func (s *stream) Send(msg interface{}) error {
	if s.isSecure {
		if len(s.encKey) == 0 {
			return ErrEmptyEncryptionKey
		}
		msg = codec.NewPacket(msg, s.encKey)
	}
	switch s.typ {
	case ServerStream:
		return s.srvSend(msg)
	case ClientStream:
		return s.clientSend(msg)
	}
	return nil
}

func (s *stream) Recv(msg interface{}) error {
	if s.isSecure {
		if len(s.encKey) == 0 {
			return ErrEmptyEncryptionKey
		}
		msg = codec.NewPacket(msg, s.encKey)
	}
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

func (s *stream) EncKey() []byte {
	return s.encKey
}

func (s *stream) FullMethod() string {
	return s.fullMethod
}

func (s *stream) MessagesReceived() int {
	return s.recvx
}
