package binarytree

import (
	"algorithm/container"
	"fmt"
	"github.com/spf13/cast"
	"strings"
)

func Serialization(head *Node) string {
	var res string
	processForSerialization(head, &res)
	return res
}

func processForSerialization(head *Node, str *string) {
	if head == nil {
		*str += "#_"
		return
	}
	*str += fmt.Sprintf("%d_", head.Value)
	processForSerialization(head.Left, str)
	processForSerialization(head.Right, str)
}

func DeSerialization(str string) *Node {
	nodeStrs := strings.Split(str, "_")
	queue := container.QueueBySlice{}
	for _, s := range nodeStrs {
		queue.Push(s)
	}
	return processForDeSerialization(&queue)
}

func processForDeSerialization(queue *container.QueueBySlice) *Node {
	if queue.Empty() {
		return nil
	}
	value := queue.Pop()
	if value.(string) == "#" {
		return nil
	}
	head := &Node{
		Value: cast.ToInt(value.(string)),
	}
	head.Left = processForDeSerialization(queue)
	head.Right = processForDeSerialization(queue)
	return head
}
