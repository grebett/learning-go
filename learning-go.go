// Mutiple files program and custom packages import
package main

import (
	"fmt"

	"github.com/grebett/stringutil" // this is a package
)

func main() {
	fmt.Printf(stringutil.Reverse(hello()))
}
