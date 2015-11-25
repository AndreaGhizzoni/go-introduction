// Read all the content of file given as path with -path flag with os and io
// primitive. Using an efficient way to concatenate strings (String buffer)
//
//author: Andrea Ghizzoni
//info:   andrea.ghz@gmail.com
package main

import (
    "fmt"
    "io"
    "os"
    "flag"
    "strings"
    "bytes"
)

func readFile( fileName string ) string {
    f, err := os.Open( fileName ); defer f.Close()
    if err != nil{
        panic(err)
    }

    var content bytes.Buffer // this is a string buffer
    buf := make( []byte, 1024 )
    for {
        n, err := f.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 { break }
        content.Write(buf)
    }

    return content.String()
}

func main(){
    path := flag.String("path", "", "the absolute paht of file to read" )
    flag.Parse()

    if strings.EqualFold( *path, "" ){
        fmt.Printf( "-path=/path/to/file not found" )
        return
    }

    fmt.Printf( "File content:\n%s", readFile(*path) )
}

