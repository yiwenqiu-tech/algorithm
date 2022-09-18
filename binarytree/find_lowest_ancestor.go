package binarytree

import "algorithm/container"

// 给定两个二叉树的节点node1和node2, 找到他们的最低公共祖先节点

func FindLowestAncestor1(head *Node, node1 *Node, node2 *Node) *Node {
	if head == nil {
		return nil
	}
	if node1 == node2 {
		return node1
	}
	fatherNodeMap := make(map[*Node]*Node) // value：存储当前节点，key：存储父节点
	nodeQueue := container.QueueBySlice{}
	nodeQueue.Push(head)
	fatherNodeMap[head] = head // 头节点特殊处理,value为自己
	for !nodeQueue.Empty() {
		node := nodeQueue.Pop().(*Node)
		if node.Left != nil {
			nodeQueue.Push(node.Left)
			fatherNodeMap[node.Left] = node
		}
		if node.Right != nil {
			nodeQueue.Push(node.Right)
			fatherNodeMap[node.Right] = node
		}
	}
	node1List := make(map[*Node]struct{})
	var cur = node1
	for cur != fatherNodeMap[cur] {
		node1List[cur] = struct{}{}
		cur = fatherNodeMap[cur]
	}
	cur = node2
	for cur != fatherNodeMap[cur] {
		if _, ok := node1List[cur]; ok {
			return cur
		}
		cur = fatherNodeMap[cur]
	}
	// 如果node1、node2都在同一颗二叉树上，一般不会出现该问题
	return nil
}

// 最低公共祖先，分情况讨论：
//	1.node1与node2互为父子节点
// 	2.node1与node2分别处于父节点的两侧
func FindLowestAncestor2(head *Node, node1 *Node, node2 *Node) *Node {
	if head == nil || head == node1 || head == node2 {
		return head
	}
	left := FindLowestAncestor2(head.Left, node1, node2)
	right := FindLowestAncestor2(head.Right, node1, node2)
	if left != nil && right != nil {
		return head
	}
	if left != nil {
		return left
	}
	return right
}
