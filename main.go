// Working with interfaces
package main

import (
	"fmt"
)

type Shaper interface {
	Area() int
}

type Square struct {
	side int
}

func (s *Square) Area() int {
	return s.side * s.side
}

func main() {
	s := new(Square)
	s.side = 5

	// check if interface is implemented
	var areaIntf Shaper
	areaIntf = s

	// use method from the interface variable
	fmt.Println(areaIntf.Area())
}
