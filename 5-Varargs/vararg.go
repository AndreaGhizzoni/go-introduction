//Basic example of var args function
package main

import "fmt"

func Min(args ...int) int {
    min := int(^uint(0)>>1)  // largest possible int
    for _, x := range args { // args has type []int 
        if min > x { min = x }
    }
    return min
}

func main(){
    fmt.Println( Min(1, 2, 3), Min(-27), Min(), Min( 7, 1, 5) )
}

