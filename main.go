// Reading and writing
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter something:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The input entered is: %s", input)
	}
}
