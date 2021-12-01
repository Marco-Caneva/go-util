package main

/*      v---   oggetto ricevente o target di chiamata del metodo (non func) */

func (first *Node) lenght() int {
	c := 0
	for first != nil {
		c++
		first = first.Next
	}
	return c
}
