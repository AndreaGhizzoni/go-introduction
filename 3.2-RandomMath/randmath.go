//generates random numbers from math/rand module
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.New(rand.NewSource(99))
	fmt.Printf("Int: %d\n", r.Intn(10))
	fmt.Printf("Perm: %v\n", r.Perm(10))
}
