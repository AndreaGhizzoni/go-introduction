// Union of example 1 and 2
package main

import (
    "fmt"
    "flag"
    "strings"
    "io/ioutil"
    "os"
)

func main(){
    path := flag.String("path", "", "the absolute path of file to read")
    flag.Parse()

    if strings.EqualFold( *path, "" ) {
        fmt.Println( "absolute path not specified" )
        os.Exit(1)
    }

    out, err := ioutil.ReadFile(*path)
    if err != nil {
        fmt.Println( err )
        os.Exit(1)
    }

    fmt.Printf( string(out) )
}

