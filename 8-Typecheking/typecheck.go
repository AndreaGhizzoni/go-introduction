// Basic type checking
package main

import "fmt"

// create a new Element of type interface
type Element interface {}

// create a vector of Element with methods
type Vector []Element
var pos = 0

// returns the i'th element
func (p *Vector) At(i int) Element {
    return (*p)[i]
}

// set always the 0'th element
func (p *Vector) Set(e Element) {
    (*p)[pos] = e
    pos += 1
}

//return the size of Vector
func (p *Vector) Size() int {
    return pos
}

func main(){
    v := make(Vector, 10)
    v.Set(10)
    v.Set(10.0)
    v.Set("asd")

    for i:=0; i<v.Size(); i++ {
        elem := v.At(i)
        switch v := elem.(type) {
        case int:
            fmt.Printf("v[%d] = %d, type: int\n", i, v)
        case float64:
            fmt.Printf("v[%d] = %f, type: float\n", i, v)
        default:
            fmt.Printf("v[%d] is unknown type\n", i)
        }
    }
}

