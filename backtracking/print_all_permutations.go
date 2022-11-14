package backtracking

// 打印一个字符串的全部排列

// 进阶：打印一个字符串的全部排列，要求不要出现重复的排列

func printAllPermutations(input string) []string {
	var res []string
	processForPrintAllPermutations([]byte(input), 0, &res)
	return res
}

func processForPrintAllPermutations(chars []byte, i int, res *[]string) {
	if i == len(chars) {
		*res = append(*res, string(chars))
		return
	}
	for j := i; j < len(chars); j++ {
		swap(chars, i, j)
		processForPrintAllPermutations(chars, i+1, res) // 进入下一层递归
		swap(chars, i, j)                               // 需要交换回来，进行下一轮交互的尝试
	}

}

func swap(chars []byte, i, j int) {
	tmp := chars[i]
	chars[i] = chars[j]
	chars[j] = tmp
}

func printAllPermutationsWithoutDup(input string) []string {
	var res []string
	processForPrintAllPermutationsWithoutDup([]byte(input), 0, &res)
	return res
}

func processForPrintAllPermutationsWithoutDup(chars []byte, i int, res *[]string) {
	if i == len(chars) {
		*res = append(*res, string(chars))
		return
	}
	tried := make([]bool, 26) // 假设字符串中只有26个小写字母
	for j := i; j < len(chars); j++ {
		tmp := chars[j] - 'a' // 找出对应放在数组的位置; 分支限界法。
		if tried[tmp] {       // 如果之前已经try过, 直接continue
			continue
		}
		tried[tmp] = true
		swap(chars, i, j)
		processForPrintAllPermutationsWithoutDup(chars, i+1, res) // 进入下一层递归
		swap(chars, i, j)                                         // 需要交换回来，进行下一轮交互的尝试
	}

}
