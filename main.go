// Reading and writing
package main

import (
	"fmt"
)

func main() {
	var (
		name, animal string
	)

	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Printf("Your name is %s\n", name)

	fmt.Println("And what is your favorite animal?")
	fmt.Scanf("the %s", &animal)
	fmt.Printf("Ah, it's the %s\n", animal)
}
