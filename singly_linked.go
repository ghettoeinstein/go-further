package main

import "fmt"

type Node struct {
	Value int 
	Next *Node
}


func main() {
	aa := Node{Value: 1}
	bb := Node{Value: 2}
	cc := Node{Value: 3}
	
	aa.Next = &bb
	bb.Next = &cc

	Print(&aa)
	Print(&bb)
	// Lets walk the the linked list

	nodeCount := Len(&aa)
	fmt.Println("Node count is: ", nodeCount)

	// Assign a new node to the beginning of the linked list
	zz := &Node{Value: 0}
	zz.Next = &aa
	
	// Update the nodeCount variabble to the new linked list length
	nodeCount = Len(zz)
	fmt.Println("Node count is: ", nodeCount)

	Print(zz)
}


func Print(root *Node) {
	nodeWalk := root
	for nodeWalk.Next != nil {
        fmt.Println(nodeWalk.Value)
        nodeWalk = nodeWalk.Next
    }
}

func Len(root *Node) int {
	nodeWalk := root
    count := 1
    for nodeWalk.Next != nil {
        count++
        nodeWalk = nodeWalk.Next
    }
    return count
}