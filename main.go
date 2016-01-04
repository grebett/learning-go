// Working with structs
package main

import (
	"fmt"
)

type Dog struct {
	name string
}

func (d *Dog) bark() {
	fmt.Printf("%s barks: woof woof!\n", d.name)
}

func main() {
	var bobby *Dog

	bobby = &Dog{"bobby"}
	bobby.bark()
}
