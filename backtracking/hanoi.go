package backtracking

import "fmt"

// 汉诺塔问题
// 	打印n层汉诺塔从最左边移动到最右边的全部过程

func hanoi(n int) {
	processForHanoi(3, "左", "右", "中")
}

func processForHanoi(i int, from, to, other string) {
	if i == 1 {
		fmt.Printf("move %v from %s to %s\n", i, from, to)
	} else {
		processForHanoi(i-1, from, other, to)              // 子问题：把i-1层挪到other
		fmt.Printf("move %v from %s to %s\n", i, from, to) // 挪第i层
		processForHanoi(i-1, other, to, from)              // 把i-1从other挪到to上
	}
}
