package main

import "fmt"

type Coordinate struct {
	x, y int
}

func main() {  
	var c *Coordinate = new(Coordinate)
	// or var c *Coordinate = Coordinate(3,4)
	c.x = 1
	c.y = 2
}