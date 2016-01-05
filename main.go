// Error handling
package main

import (
	"fmt"
	"time"
)

// creating a custom struct
type CustomError struct {
	when    time.Time
	message string
}

// using this struct as an error - i.e adding an Error method to the struct
func (e *CustomError) Error() string {
	return fmt.Sprintf("Error (%v): %s", e.when, e.message)
}

func Test(i int) (int, error) {
	if i < 0 {
		// calling the error method immediatly using a struct pointer
		return 0, &CustomError{time.Now(), "the i number should not be negative"}
	} else {
		return i, nil
	}
}

func main() {
	fmt.Println(Test(-1))
	fmt.Println(Test(1))
}
