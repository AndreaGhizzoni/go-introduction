//Example of Go lang program to calculate n-Fibonacci number.
package main

import (
    "flag"
    "fmt"
)

func Fibonacci( n int ) (int64) {
    var f, s, t int64
    f, s, t = 1, 1, 1
    for i:=2; i<n; i++ {
        f = s
        s = t
        t = f+s
    }
    return t
}

func main(){
    p := flag.Int("of", -1, "the n-Fibonacci number to calculate")
    flag.Parse()

    n := (*p)
    if n == -1 {
        fmt.Println("-of=<n> not found")
        return
    }

    fmt.Printf("Fibonacci(%v) = %v\n", n, Fibonacci(n) )
}

