package main

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	cloud "gRPC/chat"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, in *cloud.Message) (*cloud.Message, error) {
	fmt.Println(in.Body)
	return &cloud.Message{Body: "beep boop i am a server"}, nil
}

//gets called when a file is uploaded to the server
func (s *Server) Upload(ctx context.Context, in *cloud.File) (*cloud.Message, error) {
	//create a local file with the bytes of the request
	err := ioutil.WriteFile(in.Filename, in.File, 644)
	if err != nil {
		return nil, errors.New("error writing file to server")
	}
	return &cloud.Message{Body: "File received successfully"}, nil
}

//gets called when a client wants to download a file from the server
func (s *Server) Download(ctx context.Context, in *cloud.Message) (*cloud.File, error) {
	fmt.Println("lookin for file ", in.Body)
	b, err := ioutil.ReadFile(in.Body) // b has type []byte
	if err != nil {
		log.Fatal(err)
	}
	return &cloud.File{
		File:     b,
		Filename: in.Body,
	}, nil
}

func (s *Server) GetFiles(ctx context.Context, in *cloud.Message) (*cloud.Message, error) {
	foldername := in.Body
	myMap := make(map[string]string)
	err := filepath.Walk(foldername, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			myMap[info.Name()] = "folder"
		} else {
			myMap[info.Name()] = "file"
		}
		return nil
	})
	if err != nil {
		return nil, errors.New(fmt.Sprint("had a problem wile tryint to read all files in folder", foldername))
	}
	jsonString, _ := json.Marshal(myMap)
	fmt.Println(jsonString)

	return &cloud.Message{
		Body: b64.StdEncoding.EncodeToString(jsonString),
	}, nil
}
