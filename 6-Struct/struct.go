//Simple structure with method
package main

import "fmt"

type Point struct {
	x, y float64
}

func (p *Point) Sum() float64 {
	return p.x + p.y
}

func main() {
	p := &Point{1, 2}
	fmt.Printf("p.Sum() = %f\n", p.Sum())

	p1 := &Point{x: 1, y: 2}
	fmt.Printf("p1.Sum() = %f\n", p1.Sum())

	p2 := new(Point)
	p2.x = 1
	p2.y = 1
	fmt.Printf("p2.Sum() = %f\n", p2.Sum())
}
