package list

import "algorithm/sort"

//将单向链表按某值划分成左边小、中间相等、右边大的形式
//【题目】
//给定一个单链表的头节点head，节点的值类型是整型，再给定一个整数pivot。实现一个调整链表的函数，将链表调整为左部分都是值小于pivot的节点，
//中间部分都是值等于pivot的节点，右部分都是值大于pivot的节点。
//
//【进阶】在实现原问题功能的基础上增加如下的要求
//【要求】调整后所有小于pivot的节点之间的相对顺序和调整前一样
//【要求】调整后所有等于pivot的节点之间的相对顺序和调整前一样
//【要求】调整后所有大于pivot的节点之间的相对顺序和调整前一样
//【要求】时间复杂度请达到O(N)，额外空间复杂度请达到O(1)。

// 通过转换数组，再通过数组partition处理，最后再串起处理好的数组；时间复杂度O(N), 空间复杂度O(N)
func PartitionByArray(head *SingleLinkedNode, num int) *SingleLinkedNode {
	var nodeValues []int
	for head != nil {
		nodeValues = append(nodeValues, head.Value)
		head = head.Next
	}
	res := sort.DutchNationalFlag(nodeValues, num)
	var tmp = &SingleLinkedNode{}
	var last = tmp
	for _, value := range res {
		last.Next = &SingleLinkedNode{
			Value: value,
		}
		last = last.Next
	}
	return tmp.Next
}

// 定义小于区、等于区、大于区的头尾指针，遍历单向链表，依次放置不同区域，最后串起来；时间复杂度O(N), 空间复杂度O(a)
func PartitionByPointer(head *SingleLinkedNode, num int) *SingleLinkedNode {
	// 定义六个变量，小于区的头指针，小于区的尾指针，等于区的头指针，等于区的尾指针，大于区的头指针，大于区的尾指针
	var lh, lt, eh, et, mh, mt *SingleLinkedNode
	var curNode = head
	for curNode != nil {
		next := curNode.Next
		curNode.Next = nil // 需要把下一个指针清空，避免尾指针仍存在原next信息
		if curNode.Value < num {
			if lh == nil || lt == nil {
				lh = curNode
				lt = curNode
			} else {
				lt.Next = curNode
				lt = curNode
			}
		} else if curNode.Value == num {
			if eh == nil || et == nil {
				eh = curNode
				et = curNode
			} else {
				et.Next = curNode
				et = curNode
			}
		} else {
			if mh == nil || mt == nil {
				mh = curNode
				mt = curNode
			} else {
				mt.Next = curNode
				mt = curNode
			}
		}
		curNode = next
	}
	if lt != nil {
		if et != nil {
			lt.Next = eh
			et.Next = mh
		} else {
			lt.Next = mh
		}
		return lh
	} else {
		if et != nil {
			et.Next = mh
			return eh
		} else {
			return mh
		}
	}
}
