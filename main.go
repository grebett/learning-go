// Working with arrays – vertices
package main

import (
	"fmt"
)

type vertex [3]int

func main() {
	a := vertex{25, -12, 4}

	fmt.Println(invert(a))
}

func invert(v vertex) vertex {
	for i := range v {
		v[i] *= -1
	}
	return v
}
