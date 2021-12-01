/* definisco il tipo di base dei nodi che fanno parte di una lista concatenata*/
package main

type Node struct {
	Value string
	Next *Node
}