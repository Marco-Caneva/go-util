package main

func (first *Node) AddFirst(x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.Value = x
	nuovo.Next = first

	return nuovo
}

func (first *Node) AddInOrder(x string) *Node {
	var nuovo *Node

	nuovo = new(Node)
	nuovo.Value = x
	var prev, curr *Node
	prev = nil
	curr = first

	//la condizione di uscita è se curr è nil o se value >= x

	for curr != nil && curr.Value < x {
		prev = curr
		curr = curr.Next
	}
	if prev == nil {
		return first.AddFirst(x)
	} else {
		prev.Next = nuovo
		nuovo.Next = curr
		return first
	}
}
