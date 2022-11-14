package backtracking

// 规定1和A对应、2和B对应、3和C对应...
//	那么一个数字字符串比如"111"，就可以转化为"AAA"、"KA"和"AK"。
//	给定一个只有数字字符组成的字符串str，返回有多少种转化结果。

func convertToLetterString(input string) int {
	return processForConvertToLetterString([]byte(input), 0)
}

// 解体思路： i之前的位置，如何转化已经做过决定了。
// 	看i... 有多少种转化的结果。
func processForConvertToLetterString(bs []byte, i int) int {
	if i == len(bs) {
		return 1
	}

	if bs[i] == '0' {
		return 0
	}

	if bs[i] == '1' {
		res := processForConvertToLetterString(bs, i+1) // i自己作为单独的部分，后面有多少种
		if i+1 < len(bs) {
			res += processForConvertToLetterString(bs, i+2) // i和i+1作为共同部分，后面有多少种
		}
		return res
	}

	if bs[i] == '2' {
		res := processForConvertToLetterString(bs, i+1) // i自己作为单独的部分，后面有多少种
		if i+1 < len(bs) && bs[i+1] < '6' {
			res += processForConvertToLetterString(bs, i+2) // i和i+1作为共同部分，后面有多少种
		}
		return res
	}
	return processForConvertToLetterString(bs, i+1)
}
