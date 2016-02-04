// example of ranging over closing channel
package main

import (
	"fmt"
	"time"
)

func doWork(queue chan<- string) {
	queue <- "Hello"
	time.Sleep(time.Second)
	queue <- "World"
	close(queue)
}

func main() {
	// make a channel of strings
	queue := make(chan string)

	// run the go-routine
	go doWork(queue)

	// we can use a range statement to retrieve every string from the channel
	// when they are put in it. This can be used if and only if the channel it's
	// closed at some point by go-routine, otherwise it will no stop retrieving
	for elem := range queue {
		fmt.Println(elem)
	}
}
