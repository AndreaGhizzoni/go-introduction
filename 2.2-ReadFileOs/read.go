// Read all the content of file given as path with -path flag with os and io
// primitive
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
)

func readFile( fileName string ) (content string) {
    f, err := os.Open( fileName ); defer f.Close()
    if err != nil{
        panic(err)
    }

    buf := make( []byte, 1024 )
    for {
        n, err := f.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        if n == 0 { break }
        content += string(buf)
    }

    return
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

