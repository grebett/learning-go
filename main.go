// Error handling – panic bubbling
package main

import "fmt"

func main() {
	test()
	fmt.Println("this is executed")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("An error as been caught: %v\n", err)
		}
	}()

	fn()
	fmt.Println("this is not executed") // because panic goes to defer immediatly
}

func fn() {
	panic("oh no!")
}
