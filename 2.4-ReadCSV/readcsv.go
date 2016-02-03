// Example to read csv file with encoding/csv package and count all the entry
package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	csvFilePath := flag.String("path", "", "csv file path")
	flag.Parse()
	if strings.EqualFold(*csvFilePath, "") {
		fmt.Println("csv file path not specified. use -path <filepath>")
		return
	}

	// open the csv file
	file, err := os.Open(*csvFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// create a new csv reader based on a buffer to opened file.
	csvR := csv.NewReader(bufio.NewReader(file))

	var counter int64
	for {
		// read the entire csv entry
		_, err := csvR.Read()
		if err == io.EOF {
			break
		}
		counter++
	}
	fmt.Printf("Counted %d entry\n", counter)
}
