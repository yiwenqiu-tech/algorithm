package list

import "fmt"

// SingleLinkedNode 单向链表node
type SingleLinkedNode struct {
	Value int
	Next  *SingleLinkedNode
}

func printList(head *SingleLinkedNode) {
	for head != nil {
		fmt.Printf("%v;", head.Value)
		head = head.Next
	}
	fmt.Printf("\n")
}

// DoubleLinkedNode 双向链表node
type DoubleLinkedNode struct {
	Value int
	Next  *DoubleLinkedNode
	Last  *DoubleLinkedNode
}
