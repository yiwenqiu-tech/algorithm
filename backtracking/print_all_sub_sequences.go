package backtracking

import "fmt"

// 打印一个字符串的全部子序列，包括空字符串
// 解题思路：从左往右遍历字符串，每个字符存在要和不要两种选择

func printAllSubSequences(s string) {
	//processForPrintAllSubSeq([]byte(s), 0, []byte{})
	processForPrintAllSubSeq2([]byte(s), 0)
}

// 更省空间
func processForPrintAllSubSeq2(b []byte, i int) {
	if i == len(b) {
		fmt.Println(string(b))
		return
	}
	processForPrintAllSubSeq2(b, i+1)
	tmp := b[i]
	b[i] = byte(0) //修改为空。相当于不要某个字符
	processForPrintAllSubSeq2(b, i+1)
	b[i] = tmp // 利用系统栈特性，这里可以把tmp恢复回来。
}

func processForPrintAllSubSeq(b []byte, i int, res []byte) {
	if i == len(b) {
		fmt.Println(string(res))
		return
	}
	resKeep := copyRes(res)
	resKeep = append(resKeep, b[i])
	processForPrintAllSubSeq(b, i+1, resKeep)

	resNoKeep := copyRes(res)
	processForPrintAllSubSeq(b, i+1, resNoKeep)
}

func copyRes(res []byte) []byte {
	var output []byte
	for i := range res {
		output = append(output, res[i])
	}
	return output
}
