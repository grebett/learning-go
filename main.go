// Working with structs
package main

import (
	"fmt"
)

type Dog struct {
	name string
}

func (d *Dog) String() string {
	return "My name is " + d.name + " and I'm a dog!"
}

func main() {
	var bobby *Dog

	bobby = &Dog{"bobby"}

	fmt.Println(bobby)
}
