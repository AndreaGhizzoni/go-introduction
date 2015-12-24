//Basic hello world example.
package main

// import packages
import(
    "fmt"
    "os"
    "strings"
)

// main function. NB: no return type and no passing arguments
func main(){
    // short variable declaration
    who := "World!"

    // os.Args have access to arguments passing to main
    if len(os.Args) > 1 {
        who = strings.Join( os.Args[1:], " " )
    }

    // print on standard output
    fmt.Printf( "Hello %s\n", who );
}

