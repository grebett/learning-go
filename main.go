// Working with maps
package main

import (
	"fmt"
)

func main() {
	var m = map[string]int{"one": 1, "two": 2}
	n := make(map[string]int)

	n["three"] = 3
	n["for"] = 4

	fmt.Println(m)
	fmt.Println(n)

	if _, ok := n["five"]; ok {
		fmt.Println("`five` key exists")
	} else {
		fmt.Println("`five` key does not exist")
	}
}
