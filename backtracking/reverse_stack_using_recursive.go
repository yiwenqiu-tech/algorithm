package backtracking

// 给定一个栈，请逆序这个栈，不能申请额外的数据结构，只能使用递归函数。

import "algorithm/container"

func ReverseStackUsingRecursive(stack *container.StackBySlice) {
	if stack.Empty() {
		return
	}
	i := getStackBottom(stack)
	ReverseStackUsingRecursive(stack)
	stack.Push(i)
}

func getStackBottom(stack *container.StackBySlice) int {
	result := stack.Pop()
	if stack.Empty() {
		return result.(int)
	} else {
		last := getStackBottom(stack)
		stack.Push(result)
		return last
	}
}
