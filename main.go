// Working with slices – append
package main

import (
	"fmt"
)

func main() {
	str := "Hello world!"
	sub1 := str[:5]
	sub2 := str[5:]
	joined := append([]rune(sub1), []rune(sub2)...)
	fmt.Printf(string(joined))
}
