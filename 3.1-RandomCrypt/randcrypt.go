//Generate a series of random numbers using crypto/rand module
package main

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"fmt"
)

func main() {
	buf := make([]byte, 1024)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Errorf("Error, ", err)
		return
	}
	var nRand int64
	binary.Read(bytes.NewReader(buf), binary.LittleEndian, &nRand)
	fmt.Println(nRand)
}
