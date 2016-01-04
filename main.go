// Working with slices II
package main

import (
	"fmt"
)

func main() {
	str := "Hello world!"
	slice := []rune(str)

	for i := range slice {
		slice[i] += 1
	}
	fmt.Printf(string(slice))
}
