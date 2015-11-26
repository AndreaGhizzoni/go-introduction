// Basic file read using ioutil module
//
//author: Andrea Ghizzoni
//info:   andrea.ghz@gmail.com
package main

import (
    "fmt"
    "io/ioutil"
)

func main(){
    fileName := "/tmp/input"
    out, err := ioutil.ReadFile(fileName)
    if err != nil {
        fmt.Println( err )
        return
    }

    fmt.Printf( string(out) ) // ReadFile returns a []byte
}

