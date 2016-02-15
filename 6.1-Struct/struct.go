// Simple structure with method
package main

import "fmt"

// structure declaration
type Point struct {
	x, y float64
}

// Method associated with previous structure with (p *Point) declaration.
// p is know as "receiver" or implicit "this" parameter
func (p *Point) Sum() float64 {
	return p.x + p.y
}

func main() {
	// construct a pointer with unnamed filed
	p := &Point{1, 2}
	fmt.Printf("p.Sum() = %f\n", p.Sum())

	// construct a pointer with named filed
	p1 := &Point{x: 1, y: 2}
	fmt.Printf("p1.Sum() = %f\n", p1.Sum())

	// construct a pointer with new() built-in function
	p2 := new(Point)
	p2.x = 1
	p2.y = 1
	fmt.Printf("p2.Sum() = %f\n", p2.Sum())
}
