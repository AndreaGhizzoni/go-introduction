// go example of using a channel to get updates from long operations go routine
package main

import (
	"fmt"
	"time"
)

// function to simulate long operation
func veryLongFunction(done chan bool) {
	fmt.Print("Starting long operation...")
	time.Sleep(time.Second * 5)
	fmt.Println("done")
	done <- true
}

func main() {
	// make a channel with capacity of 1 to communicate with go-routine
	goRoutineDone := make(chan bool, 1)
	// run the go-routine as this function
	go veryLongFunction(goRoutineDone)

	// wait until go-routine put something into the channel
	<-goRoutineDone
}
