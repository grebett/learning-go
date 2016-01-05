// Reading and writing
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()

		writer := bufio.NewWriter(file)
		writer.WriteString("Hey, this is a string I wrote from a go program!")
		writer.Flush()
	}
}
