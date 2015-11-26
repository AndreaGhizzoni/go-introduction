// Simple struct with method
//
//author: Andrea Ghizzoni
//info:   andrea.ghz@gmail.com
package main

import "fmt"

type Point struct{
    x, y float64
}

func (p *Point) Abs() float64 {
    return 1
}

func main(){
    p := &Point{ 1, 2 }
    fmt.Printf("p.Abs() = %f\n", p.Abs())
}

