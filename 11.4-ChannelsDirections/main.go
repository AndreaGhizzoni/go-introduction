// example of channel directions
package main

import "fmt"

// example function of using a read-only channel ( <-chan ) and a write-only
// channel ( chan<- )
func chanCopy(from <-chan string, to chan<- string) {
	msg := <-from
	to <- msg
}

func main() {
	// make 2 channel of capacity of 1
	from := make(chan string, 1)
	to := make(chan string, 1)
	// send a string into the first channel
	from <- "Hello"
	chanCopy(from, to)
	// retrieve from the second channel
	fmt.Printf("Read from second channel: %s\n", <-to)
}
