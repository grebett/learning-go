// Error handling
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("this is a very basic string error")
	fmt.Printf("Error: %v\n", err)
}
