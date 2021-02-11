package main

import (
	cloud "gRPC/chat"
	"log"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := cloud.NewCloudClient(conn)

	client := Client{}

	client.SayHello(c)
	client.Upload(c, "upload.txt")
	client.Download(c, "download.txt")
}
