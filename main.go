// Working with structs
package main

import (
	"fmt"
)

type List struct {
	value string
	next  *List
}

func main() {
	var node *List
	var begin *List

	begin = new(List)
	begin.value = "this is the start node"
	node = addNode(begin, "and this is a second node")
	node = addNode(node, "and this is a third node")

	node = begin
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func addNode(node *List, value string) *List {
	newNode := new(List)

	newNode.value = value
	node.next = newNode

	return newNode
}
