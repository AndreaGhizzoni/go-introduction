//Read file with bufio package
package main

import(
    "fmt"
    "flag"
    "bufio"
    "strings"
    "os"
)

func main(){
    path := flag.String("path", "", "absolute path of file to read")
    flag.Parse()

    if strings.EqualFold( *path, "" ) || len(os.Args) == 1 {
        fmt.Printf("-path=/path/to/file missing\n")
        return
    }

    f, err := os.Open(*path)
    if err != nil{
        panic(err)
    }
    reader := bufio.NewReader(f)
    for {
        line, err := reader.ReadString('\n')
        fmt.Print(line)

        if err != nil{ break }
    }
}

