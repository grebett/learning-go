// goroutines
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("before helloInterval")
	go hello()
	fmt.Println("after helloInterval")
	time.Sleep(time.Second * 2)
}

func hello() {
	time.Sleep(time.Second)
	fmt.Println("hello world!")
}
