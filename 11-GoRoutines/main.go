// go routine test
package main

import "fmt"

// simple test function
func f(from string) {
	fmt.Printf("\nf() called from %s\n", from)
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println("\nf() from main function ends")
}

func main() {
	// direct call of a function
	f("main function")
	// start a new go routines that runs the function f
	go f("go routines")

	// wait until the go routine ends, then pres enter to exit
	var input string
	fmt.Scanln(&input)
}
