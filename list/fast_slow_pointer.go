package list

// 快慢指针的变种
//
//  1.
// 		a.当链表里元素数量为奇数时，当快指针走完时，慢指针停在中间元素
//  	b.当链表里元素数量为偶数时，当快指针走完时，慢指针停在对称线前一个元素
//	2.
//		a.当链表里元素数量为奇数时，当快指针走完时，慢指针停在中间元素
//      b.当链表里元素数量为偶数时，当快指针走完时，慢指针停在对称线后一个元素
//  ...

func FindNodeByFastSlowPointer1(head1 *SingleLinkedNode) *SingleLinkedNode {
	var fast = head1
	var slow = head1
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func FindNodeByFastSlowPointer2(head1 *SingleLinkedNode) *SingleLinkedNode {
	var fast = head1
	var slow = head1
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
