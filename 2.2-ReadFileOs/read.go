// Read all the content of file given as path with -path flag with os and io
// primitive. Using an efficient way to concatenate strings (String buffer)
package main

import (
    "fmt"
    "io"
    "os"
    "flag"
    "strings"
    "bytes"
)

// This Function read all the file content of given file path, and return it as
// string
func readFile( fileName string ) string {
    // use package os to open the file
    // Note defer function to close it
    f, err := os.Open( fileName ); defer f.Close()
    if err != nil{
        panic(err)
    }

    // this is a string buffer to hold the converted content
    var content bytes.Buffer
    // create a slice of byte as read buffer
    buf := make( []byte, 1024 )
    for {
        // read len(buf) bytes
        n, err := f.Read(buf)
        if err != nil && err != io.EOF {
            panic(err)
        }
        // if no bytes read, nothing else to read is left
        if n == 0 { break }
        // flush the buffer into string buffer
        content.Write(buf)
    }

    // return the content of the buffer
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

