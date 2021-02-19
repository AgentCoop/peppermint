package image_test

import (
	"context"
	"fmt"
	//"net"
	"testing"
	"google.golang.org/grpc"
	pb "github.com/AgentCoop/peppermint/internal/api/peppermint/multimedia/image"
	//"time"
)

//func UnixDialer(addr string, time time.Duration) (net.Conn, error) {
//	unixAddr, error := net.ResolveUnixAddr("unix", addr)
//}

func TestCreateGraph(T *testing.T) {
	conn, err := grpc.Dial("unix:@foo", grpc.WithInsecure())//, grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := pb.NewImageServiceClient(conn)

	req := &pb.Graph_Create_Request{Body: []byte("Hello")}
	res, reqErr := client.GraphCreate(context.Background(), req)
	if reqErr != nil {
		panic(reqErr)
	}

	fmt.Printf("res %v\n", res)

	defer conn.Close()
}
