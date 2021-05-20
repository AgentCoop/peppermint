package balancer

import "google.golang.org/grpc"

type upstream struct {

}

type downstream struct {
	ss *grpc.ServerStream
	cs *grpc.ClientStream
}

type balancer struct {

}

func withUnknowServiceHandler() grpc.StreamHandler {
	return func(srv interface{}, stream grpc.ServerStream) error {
		stream.RecvMsg()
		return nil
	}
}

func NewGrpcBalancer() *balancer {
	b := new(balancer)
	grpc := grpc.NewServer(withUnknowServiceHandler())
	return b
}
