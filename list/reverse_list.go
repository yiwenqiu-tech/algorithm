package list

import "google.golang.org/genproto/googleapis/cloud/tasks/v2"

func ReverseList(node *SingleLinkedNode) *SingleLinkedNode {
	if node == nil {
		return nil
	}
	cur := node
	var last *SingleLinkedNode = nil
	for cur != nil {
		next := cur.Next
		cur.Next = last
		last = cur
		cur = next
	}
	return last
}

func ReverseDoubleLinkedList(node *DoubleLinkedNode) *DoubleLinkedNode {
	if node == nil {
		return nil
	}
	cur := node
	for {
		temp := cur.Next
		cur.Last = cur.Next
		cur.Next = cur.Last
		if temp == nil {
			break
		}
		cur = temp
	}
	return cur
}
