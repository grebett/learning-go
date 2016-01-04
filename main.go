// Working with interfaces
package main

import (
	"fmt"
)

type Shaper interface {
	Area() int
	//	Missing() string
}

type Square struct {
	side int
}

type Rectangle struct {
	width  int
	height int
}

func (s *Square) Area() int {
	return s.side * s.side
}

func (r *Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	s := &Square{5}
	r := &Rectangle{5, 4}

	shapes := []Shaper{s, r}
	for _, shape := range shapes {
		fmt.Println(shape)
		fmt.Println(shape.Area())
	}
}
