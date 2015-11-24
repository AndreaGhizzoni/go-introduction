// Introduction to flag module
//
//author: Andrea Ghizzoni
//info:   andrea.ghz@gmail.com
package main

import (
	"flag"
	"fmt"
)

func main() {
	stringFlag := flag.String("str", "noStr", "a string flag value")
	boolFlag := flag.Bool("bool", false, "a boolean flag value")
	intFlag := flag.Int("int", -1, "a int flag value")
	flag.Parse()

	fmt.Printf("str: %s\n", *stringFlag)
	fmt.Printf("bool: %t\n", *boolFlag)
	fmt.Printf("int: %d\n", *intFlag)
	fmt.Printf("the rest of Flags: %s\n", flag.Args())
}
