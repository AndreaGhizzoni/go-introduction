// Union of example 1 and 2
//
//author: Andrea Ghizzoni
//info:   andrea.ghz@gmail.com
package main

import (
    "fmt"
    "flag"
    "strings"
    "io/ioutil"
)

func main(){
    path := flag.String("path", "", "the absolute path of file to read")
    flag.Parse()

    if strings.EqualFold( *path, "" ) {
        fmt.Println( "absolute path not specified" )
        return
    }

    out, err := ioutil.ReadFile(*path)
    if err != nil {
        fmt.Println( err )
        return
    }

    fmt.Printf( string(out) ) // ReadFile returns a []byte
}

