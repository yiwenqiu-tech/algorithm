package binarytree

import (
	"algorithm/container"
	"math"
)

// BFS, breadth first search（宽度优先遍历（广度优先遍历、层次遍历））

func BFS(head *Node) []int {
	if head == nil {
		return nil
	}
	queue := container.QueueBySlice{}
	queue.Push(head)
	var res []int
	for !queue.Empty() {
		value := queue.Pop().(*Node)
		res = append(res, value.Value)
		if value.Left != nil {
			queue.Push(value.Left)
		}
		if value.Right != nil {
			queue.Push(value.Right)
		}
	}
	return res
}

func FindMaxWidthUseMap(head *Node) int {
	if head == nil {
		return 0
	}
	queue := container.QueueBySlice{}
	queue.Push(head)
	// 记录当前节点层次
	var levelMap = make(map[*Node]int)
	levelMap[head] = 1
	// 记录当前遍历到哪一层
	currentLevel := 1
	// 记录当前遍历层的宽度
	currentNum := 0
	// 记录最大宽度
	max := math.MinInt32
	for !queue.Empty() {
		value := queue.Pop().(*Node)
		level := levelMap[value]
		if level == currentLevel {
			currentNum++
		} else {
			// level != currentLevel, 代表上一层已经结束遍历，开始结算
			// 	用上一层统计的currentNum与max对比，若大于max则记录currentNum
			if currentNum > max {
				max = currentNum
			}
			// 重置currentNum与currentLevel
			currentNum = 1
			currentLevel = level
		}
		if value.Left != nil {
			queue.Push(value.Left)
			levelMap[value.Left] = level + 1
		}
		if value.Right != nil {
			queue.Push(value.Right)
			levelMap[value.Right] = level + 1
		}
	}
	// 最后一轮更新max
	if currentNum > max {
		max = currentNum
	}
	return max
}

func FindMaxWidthNoUseMap(head *Node) int {
	if head == nil {
		return 0
	}
	queue := container.QueueBySlice{}
	queue.Push(head)

	curLevelEnd := head
	var nextLevelEnd *Node = nil
	curLevelNum := 0
	max := math.MinInt32

	for !queue.Empty() {
		curNode := queue.Pop().(*Node)
		curLevelNum++
		if curNode.Left != nil {
			queue.Push(curNode.Left)
			nextLevelEnd = curNode.Left
		}
		if curNode.Right != nil {
			queue.Push(curNode.Right)
			nextLevelEnd = curNode.Right
		}

		// 当前遍历到当前层最后一个元素了，则结算本层
		if curLevelEnd == curNode {
			if curLevelNum > max {
				max = curLevelNum
			}
			curLevelNum = 0
			curLevelEnd = nextLevelEnd
			nextLevelEnd = nil
		}
	}
	return max
}
