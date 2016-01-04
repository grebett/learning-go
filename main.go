// Working with structs
package main

import (
	"fmt"
)

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p := &Person{"Grégory", "Betton", 30}
	q := new(Person)

	q.firstName = "Marc"
	q.lastName = "Twain"
	q.age = 186
	fmt.Printf("%v and %v", p, q)
}
