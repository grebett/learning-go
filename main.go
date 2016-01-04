// Working with arrays – vertices by reference
package main

import (
	"fmt"
)

type vertex [3]int

func main() {
	a := vertex{25, -12, 4}
	invert(&a)
	fmt.Println(a)
}

func invert(v *vertex) {
	for i := range v {
		v[i] *= -1
	}
}
