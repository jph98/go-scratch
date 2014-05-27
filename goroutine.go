package main

import "fmt"

func main() {

    go func(msg string) {
        fmt.Println(msg)
    }("going, press enter to exit")
    
    var input string
    fmt.Scanln(&input)
    fmt.Println("done")
}