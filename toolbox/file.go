package toolbox

import (
	"fmt"
	"os"
	"path/filepath"
)

type MyFile struct{}

//Gets a filepath and returns the filetype as string
func (f *MyFile) GetFileType(path string) string {
	fileinfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("can't find file")
	}
	filename := fileinfo.Name()
	filetype := filepath.Ext(filename)
	return filetype
}
func (f *MyFile) GetFileName(path string) string {
	fileinfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("can't find file")
	}
	filename := fileinfo.Name()
	return filename
}
