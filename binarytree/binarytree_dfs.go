package binarytree

// DFS-深度优先遍历，二叉树深度优先遍历分别 先序遍历、中序遍历、后序遍历。
// 	先序遍历：头左右
//  中序遍历：左头右
//	后序遍历：左右头

import (
	"algorithm/container"
	"fmt"
)

// RecursionOrder 递归序
func RecursionOrder(head *Node, res *[]int) {
	if head == nil {
		return
	}
	fmt.Printf("%d, ", head.Value)
	*res = append(*res, head.Value)
	RecursionOrder(head.Left, res)
	fmt.Printf("%d, ", head.Value)
	*res = append(*res, head.Value)
	RecursionOrder(head.Right, res)
	fmt.Printf("%d, ", head.Value)
	*res = append(*res, head.Value)
}

// PreOrder 先序遍历
func PreOrder(head *Node, res *[]int) {
	if head == nil {
		return
	}
	*res = append(*res, head.Value)
	PreOrder(head.Left, res)
	PreOrder(head.Right, res)
}

// MidOrder 中序遍历
func MidOrder(head *Node, res *[]int) {
	if head == nil {
		return
	}
	MidOrder(head.Left, res)
	*res = append(*res, head.Value)
	MidOrder(head.Right, res)
}

// PosOrder 后序遍历
func PosOrder(head *Node, res *[]int) {
	if head == nil {
		return
	}
	PosOrder(head.Left, res)
	PosOrder(head.Right, res)
	*res = append(*res, head.Value)
}

// PreOrderUnRecursive 先序遍历（非递归）
func PreOrderUnRecursive(head *Node, res *[]int) {
	if head == nil {
		return
	}
	stack := container.StackBySlice{}
	stack.Push(head)
	for !stack.Empty() {
		node := stack.Pop().(*Node)
		*res = append(*res, node.Value)
		if node.Right != nil {
			stack.Push(node.Right)
		}
		if node.Left != nil {
			stack.Push(node.Left)
		}
	}
}

// MidOrderUnRecursive 中序遍历（非递归）
func MidOrderUnRecursive(head *Node, res *[]int) {
	if head == nil {
		return
	}
	tmp := head
	stack := container.StackBySlice{}
	for {
		if tmp != nil { // 优先把左侧压栈
			stack.Push(tmp)
			tmp = tmp.Left
		} else { // 当左节点全压进去后，开始出栈并打印，周而复始右节点
			if stack.Empty() {
				break
			}
			tmp = stack.Pop().(*Node)
			*res = append(*res, tmp.Value)
			tmp = tmp.Right
		}
	}
}

// PosOrderUnRecursive 后序遍历（非递归）
func PosOrderUnRecursive(head *Node, res *[]int) {
	if head == nil {
		return
	}
	stack := container.StackBySlice{}
	stack.Push(head)
	resStack := container.StackBySlice{}
	for !stack.Empty() {
		node := stack.Pop().(*Node)
		resStack.Push(node)
		if node.Left != nil {
			stack.Push(node.Left)
		}
		if node.Right != nil {
			stack.Push(node.Right)
		}
	}
	for !resStack.Empty() {
		node := resStack.Pop().(*Node)
		*res = append(*res, node.Value)
	}
}
