//Simple use of Interface
package main

import "fmt"

// interface declaration
type SumInterface interface {
	Sum() int
}

// structure declaration that implements SumInterface
type Point struct {
	x, y int
}

func (p *Point) Sum() int {
	return p.x + p.y
}

func main() {
	var si SumInterface
	si = &Point{1, 2}
	fmt.Printf("SumInterface.Sum = %d\n", si.Sum())
}
