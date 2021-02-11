package main

import (
	"fmt"
	cloud "gRPC/chat"
	"gRPC/toolbox"
	"io/ioutil"
	"log"

	"golang.org/x/net/context"
)

type Client struct{}

//SayHello sends a message to the server and prints the response
func (client *Client) SayHello(c cloud.CloudClient) {
	response, err := c.SayHello(context.Background(), &cloud.Message{Body: "Hello From Client!"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response.Body)
}

//uploads the file from path to the cloud
func (client *Client) Upload(c cloud.CloudClient, path string) {
	//reading a local file to upload to the server
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("can't find file")
	}
	//initialise toolbox struct myfile so i can get filename
	myfile := toolbox.MyFile{}
	filename := myfile.GetFileName(path)
	//send request to the server
	response, err := c.Upload(context.Background(), &cloud.File{File: file, Filename: filename})
	if err != nil {
		log.Fatal("Error when calling Upload")
	}
	fmt.Println(response.Body)
}

//downloads the file defined in name from the cloud
func (client *Client) Download(c cloud.CloudClient, filename string) {
	response, err := c.Download(context.Background(), &cloud.Message{Body: filename})
	if err != nil {
		log.Fatal("Error when receiving file from cloud")
	}
	err = ioutil.WriteFile(filename, response.File, 644)
	if err != nil {
		log.Fatal("Error when writing file onto client")
	}
}
