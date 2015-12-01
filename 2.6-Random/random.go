//combined 2.4 and 2.5 example
package main

import (
    "fmt"
    crand "crypto/rand"
    "encoding/binary"
    mrand "math/rand"
    "bytes"
)

// compute the seed
func makeSeed() (seed int64) {
    buff := make([]byte, 1024)
    _, err := crand.Read(buff)
    if err != nil {
        fmt.Errorf("Error, ", err)
        return
    }
    binary.Read(bytes.NewReader(buff), binary.LittleEndian, &seed)
    return
}

func main(){
    seed := makeSeed()
    r := mrand.New(mrand.NewSource(seed))
    fmt.Printf("Seed: %v\n", seed)
    fmt.Printf("r.Int63(): %v\n", r.Int63())
}

