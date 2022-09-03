package list

//【题目】给定两个可能有环也可能无环的单链表，头节点head1和head2。请实现一个函数，如果两个链表相交，请返回相交的第一个节点。如果不相交，返回null
//【要求】如果两个链表长度之和为N，时间复杂度请达到O(N)，额外空间复杂度请达到O(1)。

// 题目分析，题目里提到两个可能有环，可能没环的单链表，所以需要分开讨论，先判断链表有无环，然后根据有无环在分别处理

func FindFirstIntersectNode(head1 *SingleLinkedNode, head2 *SingleLinkedNode) *SingleLinkedNode {
	if head1 == nil || head2 == nil { // 处理边界情况
		return nil
	}
	cycleEntranceForList1 := FindListCycleEntrance(head1)
	cycleEntranceForList2 := FindListCycleEntrance(head2)
	if cycleEntranceForList1 == nil && cycleEntranceForList2 == nil {
		return processForNoCycle(head1, head2)
	} else if (cycleEntranceForList2 != nil && cycleEntranceForList1 == nil) ||
		(cycleEntranceForList1 != nil && cycleEntranceForList2 == nil) {
		// 一个有环，一个没环，肯定不会相交
		return nil
	} else {
		return processForHaveCycle(head1, head2, cycleEntranceForList1, cycleEntranceForList2)
	}
}

func processForHaveCycle(head1 *SingleLinkedNode, head2 *SingleLinkedNode,
	entrance1 *SingleLinkedNode, entrance2 *SingleLinkedNode) *SingleLinkedNode {
	// 都有环的情况，存在三种情况，1.各自成环、2.环外相交、3.环内相交
	if entrance1 == entrance2 {
		// 环外相交情况，采用无环链表的方式求解
		head1ToEntrance1Len := 0
		tmp1 := head1
		for tmp1 != entrance1 {
			head1ToEntrance1Len++
			if tmp1.Next == entrance1 {
				break
			}
			tmp1 = tmp1.Next
		}
		head2ToEntrance2Len := 0
		tmp2 := head2
		for tmp2 != entrance2 {
			head2ToEntrance2Len++
			if tmp2.Next == entrance2 {
				break
			}
			tmp2 = tmp2.Next
		}
		tmp1 = head1
		tmp2 = head2
		if head1ToEntrance1Len > head2ToEntrance2Len {
			tmp1 = head1
			diff := head1ToEntrance1Len - head2ToEntrance2Len
			for diff > 0 {
				tmp1 = tmp1.Next
				diff--
			}
		} else {
			tmp2 = head2
			diff := head2ToEntrance2Len - head1ToEntrance1Len
			for diff > 0 {
				tmp2 = tmp2.Next
				diff--
			}
		}
		for {
			if tmp1 == tmp2 {
				return tmp1
			}
			tmp1 = tmp1.Next
			tmp2 = tmp2.Next
		}
	} else {
		tmp1 := entrance1
		for {
			tmp1 = tmp1.Next
			if tmp1 == entrance2 {
				return entrance1 // 环内相交时，随便放回一个节点即可
			}
			if tmp1 == entrance1 { // 绕了一圈，还没碰到entrance2，则是两者不相交
				return nil
			}
		}
	}
}

func processForNoCycle(head1 *SingleLinkedNode, head2 *SingleLinkedNode) *SingleLinkedNode {
	// 无环链表下，若两个单链表相交，则最后的节点必定一致
	// 在此基础上，可以统计两个单链表的长度，然后让长链表先走差值步数，然后两者一起走，等遇到第一个一样的节点后，则是第一个相交点
	nodeLen1 := 0
	tmp1 := head1
	for tmp1 != nil {
		nodeLen1++
		if tmp1.Next == nil {
			break
		}
		tmp1 = tmp1.Next
	}

	nodeLen2 := 0
	tmp2 := head2
	for tmp2 != nil {
		nodeLen2++
		if tmp2.Next == nil {
			break
		}
		tmp2 = tmp2.Next
	}
	if tmp1 != tmp2 {
		return nil
	}
	tmp1 = head1
	tmp2 = head2
	if nodeLen1 > nodeLen2 {
		diff := nodeLen1 - nodeLen2
		for diff > 0 {
			tmp1 = tmp1.Next
			diff--
		}
	} else {
		diff := nodeLen2 - nodeLen1
		for diff > 0 {
			tmp2 = tmp2.Next
			diff--
		}
	}
	for {
		if tmp1 == tmp2 {
			return tmp1
		}
		tmp1 = tmp1.Next
		tmp2 = tmp2.Next
	}
}
