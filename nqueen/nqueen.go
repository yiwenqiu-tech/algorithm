package nqueen

import "math"

// N皇后问题
// N皇后问题是指在N*N的棋盘上要摆N个皇后，要求任何两个皇后不同行、不同列，也不在同一条斜线上。
//	给定一个整数n，返回n皇后的摆法有多少种。
//	n=1，返回1。
//	n=2或3，2皇后和3皇后问题无论怎么摆都不行，返回0。n=8，返回92。

func NQueen(n int) int {
	if n < 1 {
		return 0
	}
	var record = make([]int, n) // 记录每一行上皇后摆在哪一列
	return process(record, 0, n)
}

func process(record []int, cur int, n int) int {
	if cur == n {
		return 1
	}
	var res = 0
	for i := 0; i < n; i++ {
		if isValid(record, cur, i) {
			record[cur] = i
			res += process(record, cur+1, n)
		}
	}
	return res
}

func isValid(record []int, cur int, choose int) bool {
	if cur == 0 {
		return true
	}
	for i := 0; i < cur; i++ {
		if choose == record[i] || int(math.Abs(float64(cur-i))) == int(math.Abs(float64(choose-record[i]))) {
			return false
		}
	}
	return true
}
