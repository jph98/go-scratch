package main

import "fmt"

// You only have to specify the type once for the parameters
// Functions can return multiple results
func add(x int, y int) int {
    return x + y
}

func main() {
    fmt.Println(add(42, 13))
}