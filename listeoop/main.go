package main

import (
	"fmt"
)

func main() {
	var first *Node

	first = first.AddInOrder("ciao")
	first = first.AddInOrder("come")
	first = first.AddInOrder("stai")
	first = first.AddInOrder("aaa")
	first = first.AddInOrder("abaco")
	first = first.AddFirst("zai")
	first = first.AddInOrder("aiuto")

	fmt.Println(first.lenght())
	first.print()
}
