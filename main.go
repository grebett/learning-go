// goroutines
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(2 * time.Second)
}

func getData(ch chan string) {
	for {
		str := <-ch
		fmt.Println(str)
	}
}

func sendData(ch chan string) {
	ch <- "hello"
	time.Sleep(1 * time.Second)
	ch <- "world"
}
