//Example to read csv file
package main

import (
    "fmt"
    "flag"
    "strings"
    "os"
    "bufio"
    "encoding/csv"
    "io"
)

func main(){
    csvFilePath := flag.String("path", "", "csv file path")
    flag.Parse()
    if strings.EqualFold(*csvFilePath, "") {
        fmt.Println("csv file path not specified. use -path <filepath>")
        return
    }

    file, err := os.Open(*csvFilePath)
    if err != nil {
        fmt.Println("Error: ", err)
        return
    }

    csvR := csv.NewReader(bufio.NewReader(file))

    var counter int64
    for {
        _, err := csvR.Read()
        if err == io.EOF { break }
        counter++
    }
    fmt.Printf("Counted %d entry\n", counter)
}

