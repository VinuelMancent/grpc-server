package main

import (
	"fmt"
	cloud "gRPC/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println(lis.Addr())

	cloudServer := Server{}
	grpcServer := grpc.NewServer()

	cloud.RegisterCloudServer(grpcServer, &cloudServer)

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("err serving!")
	}
}
