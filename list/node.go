package list

// SingleLinkedNode 单向链表node
type SingleLinkedNode struct {
	Value int
	Next  *SingleLinkedNode
}

// DoubleLinkedNode 双向链表node
type DoubleLinkedNode struct {
	Value int
	Next  *DoubleLinkedNode
	Last  *DoubleLinkedNode
}
