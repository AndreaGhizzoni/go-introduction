// Basic file read using ioutil module
// see golang.org/pkg/io/ioutils
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// create a full path of file to read
	fileName := "/tmp/input"

	// use ioutil to read all the file. Note the double return parameter
	out, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	// cast to string because ioutil.ReadFile() returns a []byte
	fmt.Printf(string(out))
}
