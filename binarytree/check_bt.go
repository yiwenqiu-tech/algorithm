package binarytree

import (
	"algorithm/container"
	"math"
)

// sbt: 搜索二叉树：search binary tree，特点：左树小于根节点，根节点小于右树
func CheckSBT(head *Node) bool {
	if head == nil {
		return true
	}
	r := process(head)
	return r.sbt
}

type result struct {
	min int
	max int
	sbt bool
}

func process(head *Node) *result {
	if head == nil {
		return nil
	}
	leftResult := process(head.Left)
	rightResult := process(head.Right)
	if leftResult == nil && rightResult == nil {
		return &result{
			min: head.Value,
			max: head.Value,
			sbt: true,
		}
	}
	if leftResult == nil {
		if !rightResult.sbt {
			return &result{
				sbt: false,
			}
		}
		if rightResult.min <= head.Value {
			return &result{
				sbt: false,
			}
		}
		return &result{
			min: head.Value,
			max: rightResult.max,
			sbt: true,
		}
	}
	if rightResult == nil {
		if !leftResult.sbt {
			return &result{
				sbt: false,
			}
		}
		if leftResult.max >= head.Value {
			return &result{
				sbt: false,
			}
		}
		return &result{
			min: leftResult.min,
			max: head.Value,
			sbt: true,
		}
	}
	if !leftResult.sbt || !rightResult.sbt {
		return &result{
			sbt: false,
		}
	}
	if leftResult.max >= head.Value || rightResult.min <= head.Value {
		return &result{
			sbt: false,
		}
	}
	return &result{
		min: leftResult.min,
		max: rightResult.max,
		sbt: true,
	}
}

func CheckSBT2(head *Node) bool {
	res := process2(head)
	preMaxValue = math.MinInt32 // 重置变量
	return res
}

var preMaxValue = math.MinInt32

func process2(head *Node) bool {
	if head == nil {
		return true
	}
	leftIsBST := process2(head.Left)
	// 中序遍历操作点
	if !leftIsBST {
		return false
	}
	if head.Value <= preMaxValue {
		return false
	}
	preMaxValue = head.Value

	return process2(head.Right)
}

// CBT（Completed Binary Tree）完全二叉树
//  判断条件：
//		1.若存在某个节点只有右节点，没有左节点，则不是完全二叉树
//		2.若存在某个节点左右节点不双全，则该节点后续的节点（使用层次遍历）都是叶子节点
//  层次遍历完后，均满足上述条件，则表示是完全二叉树
func CheckCBT(head *Node) bool {
	if head == nil {
		return true
	}
	var leftNodeBegin = false // 叶子节点开始标识
	queue := container.QueueBySlice{}
	queue.Push(head)
	for !queue.Empty() {
		node := queue.Pop().(*Node)
		// 命中条件1，则直接返回false
		if node.Right != nil && node.Left == nil {
			return false
		}
		// 命中条件2，则直接返回false
		if leftNodeBegin && (node.Left != nil || node.Right != nil) {
			return false
		}
		if node.Left != nil {
			queue.Push(node.Left)
		}
		if node.Right != nil {
			queue.Push(node.Right)
		}
		if !leftNodeBegin && (node.Left == nil || node.Right == nil) {
			leftNodeBegin = true
		}
	}
	// 层次遍历完后，均满足上述条件，则表示是完全二叉树
	return true
}

// FBT (full binary tree) 满二叉树，可以用节点个数以及层数来看是否满二叉树。
// 	满二叉树：满足 2^L - 1 = N
func CheckFBT(head *Node) bool {
	res := fbtProcess(head)
	if int(math.Pow(2, float64(res.level)))-1 == res.count {
		return true
	}
	return false
}

type fbtCheckItem struct {
	level int
	count int
}

func fbtProcess(head *Node) *fbtCheckItem {
	if head == nil {
		return &fbtCheckItem{
			level: 0,
			count: 0,
		}
	}
	leftCheck := fbtProcess(head.Left)
	rightCheck := fbtProcess(head.Right)
	var curLevel = 0
	if leftCheck.level > rightCheck.level {
		curLevel = leftCheck.level + 1
	} else {
		curLevel = rightCheck.level + 1
	}
	return &fbtCheckItem{
		level: curLevel,
		count: leftCheck.count + rightCheck.count + 1,
	}
}

// BBT (Balanced binary tree) 平衡二叉树，特点：它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树
func CheckBBT(head *Node) bool {
	return bbtProcess(head).isBBT
}

type bbtCheckItem struct {
	isBBT bool
	level int
}

func bbtProcess(head *Node) *bbtCheckItem {
	if head == nil {
		return &bbtCheckItem{
			isBBT: true,
			level: 0,
		}
	}
	leftCheck := bbtProcess(head.Left)
	rightCheck := bbtProcess(head.Right)
	if !leftCheck.isBBT || !rightCheck.isBBT {
		return &bbtCheckItem{
			isBBT: false,
		}
	}
	var curLevel = 0
	if leftCheck.level > rightCheck.level {
		if leftCheck.level-rightCheck.level > 1 {
			return &bbtCheckItem{
				isBBT: false,
			}
		}
		curLevel = leftCheck.level + 1
	} else {
		if rightCheck.level-leftCheck.level > 1 {
			return &bbtCheckItem{
				isBBT: false,
			}
		}
		curLevel = rightCheck.level + 1
	}
	return &bbtCheckItem{
		isBBT: true,
		level: curLevel,
	}
}
