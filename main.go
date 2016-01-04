// Working with slices
package main

import (
	"fmt"
)

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	d := &a

	fmt.Printf("len of a: %d – cap of a: %d\n", len(a), cap(a))
	fmt.Printf("len of d: %d – cap of d: %d\n", len(d), cap(d))

	d[0] = 1
	for _, value := range a {
		fmt.Println(value)
	}
}
