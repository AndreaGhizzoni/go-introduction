// example of embedding structure
package main

import "fmt"

// simple structure
type NameObj struct {
	Name string
}

// method of NameObj
func (this *NameObj) foo() int {
	return 42
}

// type Point
type Point struct {
	X, Y float64
}

// type Square.
// NameObj is called embedded structure (or inherited structure)
type Square struct {
	NameObj       // embedded structure
	Center  Point // standard composition
	Side    float64
	Name    string //-SHADOWED
}

// method of Square.
// NOTE: without this method you can still write square.foo() but is called
// square.NameObj.foo()
func (this *Square) foo() int {
	return 24
}

func main() {
	// note that NameObj is directly initialized here, in real life is done by
	// the constructor
	square := &Square{
		NameObj: NameObj{"my-square"},
		Center:  Point{10, 10},
		Side:    10.2,
		Name:    "not-my-square",
	}

	// square.Name is called shadowed because without it will point to
	// square.NameObj.Name, the inherited structure.
	// The same trick is applied with foo() method
	fmt.Println("Accessing to shadowed field of inherited:", square.Name)
	// but is still possible to access directly to the inherited field
	fmt.Println("Accessing to inherited field:", square.NameObj.Name)
	// access to other field
	fmt.Printf("Accessing to square.Point field: %.3f %.3f\n", square.Center.X, square.Center.Y)
	fmt.Printf("Accessing to square.Side field: %.3f\n", square.Side)

	// calling foo() method
	fmt.Printf("The answer is %d not %d\n", square.NameObj.foo(), square.foo())
}
