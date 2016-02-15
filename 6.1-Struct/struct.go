// Structure with methods
package main

import (
	"errors"
	"fmt"
)

// structure declaration with two unexported field
type positivePoint struct {
	x, y float64
}

// method used as constructor of positivePoint type.
// No (p *positivePoint) delcaration needed because when this module is imported
// the syntax to create an instance of it is
// <module_name>.NewPositivePoint([...])
func NewPositivePoint(x, y float64) (*positivePoint, error) {
	if x < 0 || y < 0 {
		return nil, errors.New("positivePoint accepts positive coordinates only.")
	}
	return &positivePoint{x, y}, nil
}

// Method associated with previous structure with (p *positivePoint) declaration.
// p is know as "receiver" or implicit "this" parameter
func (p *positivePoint) Sum() float64 {
	return p.x + p.y
}

func main() {
	// note: the two ways below to construct a positivePoint are possible
	// because positivePoint is declared in this package. If was declared
	// elsewhere, these two ways can not be done because positivePoint.x and
	// positivePoint.y are unexported.

	// construct a pointer with unnamed filed
	p := &positivePoint{1, 2}
	fmt.Printf("p.Sum() = %f\n", p.Sum())

	// construct a pointer with named filed
	p1 := &positivePoint{x: 1, y: 2}
	fmt.Printf("p1.Sum() = %f\n", p1.Sum())

	// construct a pointer with new() built-in function
	p2 := new(positivePoint)
	p2.x = 1 // assignment possible because see note above.
	p2.y = 1
	fmt.Printf("p2.Sum() = %f\n", p2.Sum())

	// construct a positivePoint in the proper way
	p3, err := NewPositivePoint(10.32, 23.42)
	if err != nil {
		panic(err)
	}
	fmt.Printf("p3.Sum() = %f\n", p3.Sum())
}
