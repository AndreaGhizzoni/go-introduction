// example of go-routine timeout
package main

import (
	"fmt"
	"time"
)

func main() {
	// make a channel of string
	channel1 := make(chan string, 1)
	// create a go-routine with anonymous function
	go func() {
		time.Sleep(time.Second * 5)
		channel1 <- "Hello from go-routine"
	}()

	// select over the channel result: in case passes more then 2 seconds,
	// nothing is retrieved from channel.
	select {
	case msg := <-channel1:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Go routine timeout")
	}

	close(channel1)
}
