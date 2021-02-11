package main

import (
	"fmt"
	"gRPC/toolbox"
)

func main() {
	file := toolbox.MyFile{}
	fmt.Println("filetype:", file.GetFileType("test.txt"))
	fmt.Println("filename:", file.GetFileName("test.txt"))
}
