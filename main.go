// Working with structs
package main

import (
	"fmt"
)

type Animal struct {
	name string
	legs int
}

type Dog struct {
	Animal
}

func main() {
	bobby := new(Dog)

	bobby.name = "bobby"
	bobby.legs = 4

	fmt.Println(bobby)
}
