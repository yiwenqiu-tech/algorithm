package list

import "algorithm/container"

//【题目】给定一个单链表的头节点head，请判断该链表是否为回文结构。
//【例子】1->2->1，返回true；1->2->2->1，返回true；15->6->15，返回true；1->2->3，返回false。
//【例子】如果链表长度为N，时间复杂度达到O(N)，额外空间复杂度达到O(1)。

// CheckListPalindromeByStack 用栈实现，时间复杂度O(N)，空间复杂度O(N)
func CheckListPalindromeByStack(head *SingleLinkedNode) bool {
	stack := container.StackBySlice{}
	tmp := head
	for tmp != nil {
		stack.Push(tmp)
		tmp = tmp.Next
	}
	tmp = head
	for !stack.Empty() {
		node := stack.Pop().(*SingleLinkedNode)
		if node.Value != tmp.Value {
			return false
		}
		tmp = tmp.Next
	}
	return true
}

// CheckListPalindromeByStack2 用栈实现，时间复杂度O(N)，空间复杂度O(N)；优化版本，栈少占用n/2空间
func CheckListPalindromeByStack2(head *SingleLinkedNode) bool {
	stack := container.StackBySlice{}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	tmp := slow
	for tmp != nil {
		stack.Push(tmp)
		tmp = tmp.Next
	}
	tmp = head
	for !stack.Empty() {
		node := stack.Pop().(*SingleLinkedNode)
		if node.Value != tmp.Value {
			return false
		}
		tmp = tmp.Next
	}
	return true
}

// CheckListPalindrome 从中间节点反转比较，时间复杂度O(N)，空间复杂度O(1)
func CheckListPalindrome(head *SingleLinkedNode) bool {
	// 边界条件
	if head == nil || head.Next == nil {
		return true
	}

	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow // 中间节点
	tmp := slow.Next
	mid.Next = nil
	last := mid
	// 反转后半部分节点
	for tmp != nil {
		curNext := tmp.Next
		tmp.Next = last
		last = tmp
		tmp = curNext
	}
	tail := last // 保留尾部
	tmp1 := head // 指向头部
	tmp = tail
	res := true
	for tmp != nil && tmp1 != nil {
		if tmp1.Value != tmp.Value {
			res = false
			break
		}
		tmp1 = tmp1.Next
		tmp = tmp.Next
	}

	// 返回前重新返回回去
	last = nil
	for tail != nil {
		curNext := tail.Next
		tail.Next = last
		last = tail
		tail = curNext
	}
	return res
}
