package main

//import (
//	"fmt"
//	"context"
//	"net"
//
//	//"google.golang.org/protobuf/types/descriptorpb"
//
//	//pb "google.golang.org/protobuf"
//	"google.golang.org/grpc"
//	//"google.golang.org/protobuf/proto"
//
//	"github.com/AgentCoop/peppermint/internal/api/peppermint/multimedia/image"
//	//"github.com/AgentCoop/peppermint/internal/api/peppermint"
//)
//
//type service struct {
//	image.UnimplementedImageServiceServer
//}
//
//func (s *service) GraphCreate(ctx context.Context, req *image.Graph_Create_Request) (*image.Graph_Create_Response, error) {
//	fmt.Printf("mesg %s\n",req.Body)
//	res := new(image.Graph_Create_Response)
//	return res, nil
//}
//func (s *service) GraphAddNode(context.Context, *image.Graph_AddNode_Request) (*image.Graph_AddNode_Response, error) {
//	return nil, nil
//}
//func (s *service) GraphLoadImage(context.Context, *image.Graph_LoadImage_Request) (*image.Graph_LoadImage_Response, error) {
//	return nil, nil
//}
//func (s *service) GraphExec(context.Context, *image.Graph_Exec) (*image.Graph_AddNode_Response, error) {
//	return nil, nil
//}
//
//func main() {
//	grpc := grpc.NewServer()
//	s := &service{}
//	image.RegisterImageServiceServer(grpc, s)
//
//	unixAddr, _ := net.ResolveUnixAddr("unix", "@foo")
//	list, err := net.ListenUnix(unixAddr.Net, unixAddr)
//	if err != nil  {
//		panic(err)
//	}
//
//	grpc.Serve(list)
//
//	defer list.Close()
//}
