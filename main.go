// Error handling – panic bubbling
package main

import "fmt"

func main() {
	defer func() { fmt.Println("main defer is executed") }()
	fn1()
	fmt.Println("this is not executed")
}

func fn1() {
	defer func() { fmt.Println("fn1 defer is executed") }()
	fn2()
}

func fn2() {
	defer func() { fmt.Println("fn2 defer is executed") }()
	panic("panic in fn2")
}
