//Write a string into a given file
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	path := flag.String("path", "", "the absolute path of file to write")
	content := flag.String("content", "", "content to write")
	flag.Parse()

	if strings.EqualFold(*path, "") {
		fmt.Println("-path=/some/file/path not specified")
		return
	}

	byteContent := []byte(*content)
	err := ioutil.WriteFile(*path, byteContent, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("File %s write successfuly!\n", *path)
}
