// Example of Go lang program to calculate n-Fibonacci number.
package main

import (
	"flag"
	"fmt"
	"os"
)

// This function calculate the n-Fibonacci number.
func Fibonacci(n int64) int64 {
	var i, f, s, t int64
	f, s, t = 1, 1, 1
	for i = 2; i < n; i++ {
		f = s
		s = t
		t = f + s
	}
	return t
}

func main() {
	// take the N as main argument
	p := flag.Int64("of", -1, "the n-Fibonacci number to calculate")
	flag.Parse()

	// dereferencing pointer
	n := (*p)
	if n == -1 {
		fmt.Println("-of=<n> not found")
		os.Exit(1)
	}

	fmt.Printf("Fibonacci(%v) = %v\n", n, Fibonacci(n))
}
