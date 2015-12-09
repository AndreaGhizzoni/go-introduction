//Generate a series of random numbers using crypto/rand module
package main

import (
    "fmt"
    "crypto/rand"
    "encoding/binary"
    "bytes"
)

func main(){
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

