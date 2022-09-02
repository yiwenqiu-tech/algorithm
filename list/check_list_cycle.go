package list

// 判断单链表是否成环，若成环找出入环的第一个Node

func CheckListCycle(head *SingleLinkedNode) bool {
	if FindListCycleEntrance(head) != nil {
		return true
	}
	return false
}

// FindListCycleEntrance 寻找有环单链表的入环节点
func FindListCycleEntrance(head *SingleLinkedNode) *SingleLinkedNode {
	if head == nil {
		return nil
	}
	slow := head
	fast := head
	isCycle := false
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			isCycle = true
			break
		}
	}
	if !isCycle {
		return nil
	}
	fast = head
	for {
		if fast == slow {
			return slow
		}
		fast = fast.Next
		slow = slow.Next
	}
}
