// Simple file server
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initLogger(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func main() {
	initLogger(os.Stdout, os.Stdout, os.Stdout, os.Stderr)

	flagDir := flag.String("path", "data", "path to served directory")
	flag.Parse()
	dir := *flagDir

	if e, _ := exists(dir); !e {
		Trace.Printf("Directory does not exits. Try to create default.")
		err := os.Mkdir(dir, 0700)
		if err != nil {
			Error.Fatal("Error while creating server directory.")
		}
		Trace.Printf("Server Directory created successfully")
	}

	Trace.Printf("Server File from path: %s", dir)
	servedDir := http.Dir(dir)
	Error.Print(
		http.ListenAndServe(":8080", http.FileServer(servedDir)))
}
