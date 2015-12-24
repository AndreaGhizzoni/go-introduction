// Introduction to flag module.
// see golang.org/pkg/flag
package main

import (
	"flag"
	"fmt"
)

func main() {
    // create a flag used to pass arguments to main
	stringFlag := flag.String("str", "noStr", "a string flag value")
	boolFlag := flag.Bool("bool", false, "a boolean flag value")
	intFlag := flag.Int("int", -1, "a int flag value")

    // parse main arguments
	flag.Parse()

    // use the value contained into previous declared pointer
	fmt.Printf("str: %s\n", *stringFlag)
	fmt.Printf("bool: %t\n", *boolFlag)
	fmt.Printf("int: %d\n", *intFlag)
	fmt.Printf("the rest of Flags: %s\n", flag.Args())
}
