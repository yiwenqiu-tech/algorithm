package list

//复制含有随机指针节点的链表
//【题目】一种特殊的单链表节点类描述如下
// https://leetcode.com/problems/copy-list-with-random-pointer/
//rand指针是单链表节点结构中新增的指针，rand可能指向链表中的任意一个节点，也可能指向null。
//给定一个由Node节点类型组成的无环单链表的头节点head，请实现一个函数完成这个链表的复制，并返回复制的新链表的头节点。
//【要求】时间复杂度O(N)，额外空间复杂度O(1)

type Node struct {
	Value int
	Next  *Node
	Rand  *Node
}

// DumpListByMap 1.遍历第一遍，通过map保存节点copy，2.第二次遍历时，设置copy节点的next以及rand连接关系
func DumpListByMap(head *Node) *Node {
	var nodeMap = make(map[*Node]*Node)
	tmp := head
	// 复制节点
	for tmp != nil {
		nodeMap[tmp] = &Node{
			Value: tmp.Value,
		}
		tmp = tmp.Next
	}
	// 连接节点
	tmp = head
	for tmp != nil {
		copy := nodeMap[tmp]
		if tmp.Next != nil {
			nextCopy := nodeMap[tmp.Next]
			copy.Next = nextCopy
		}
		if tmp.Rand != nil {
			randCopy := nodeMap[tmp.Rand]
			copy.Rand = randCopy
		}
		tmp = tmp.Next
	}
	return nodeMap[head]
}

// DumpListByCopyInOriginList
func DumpListByCopyInOriginList(head *Node) *Node {
	if head == nil {
		return nil
	}
	tmp := head
	for tmp != nil {
		next := tmp.Next
		tmp.Next = &Node{
			Value: tmp.Value, // 复制一份与当前节点一致的节点
		}
		tmp.Next.Next = next
		tmp = next
	}

	tmp = head
	for tmp != nil {
		rand := tmp.Rand
		if rand != nil {
			tmp.Next.Rand = rand.Next // 设置复制节点的rand指针
		}
		tmp = tmp.Next.Next
	}

	tmp = head
	var copyHead = head.Next
	for tmp != nil {
		next := tmp.Next.Next
		copyNode := tmp.Next
		tmp.Next = next
		if next != nil {
			copyNode.Next = next.Next
		} else {
			copyNode.Next = nil
		}
		tmp = next
	}
	return copyHead
}
