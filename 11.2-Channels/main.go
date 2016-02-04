// example of go channels
package main

import "fmt"

// simple function that takes a channel of string as argument, then write a
// string in it.
func writeOnChannel(ch chan string) {
	ch <- "hello World!"
}

func main() {
	// create a channel of strings
	messages := make(chan string)
	// start a new go routine
	go writeOnChannel(messages)

	fmt.Print("Messages from go routine: ")
	// retrieve the message from go routine
	fmt.Println(<-messages)

	// closing the channel. In this way nothing could be sent into the channel
	close(messages)
}
