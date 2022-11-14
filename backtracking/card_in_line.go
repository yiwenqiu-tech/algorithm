package backtracking

// 给定一个整型数组arr，代表数值不同的纸牌排成一条线。玩家A和玩家B依次拿走每张纸牌，规定玩家A先拿，玩家B后拿，
// 但是每个玩家每次只能拿走最左或最右的纸牌，玩家A和玩家B都绝顶聪明。请返回最后获胜者的分数。
//【举例】
//arr=[1, 2, 100, 4]。
//开始时，玩家A只能拿走1或4。如果开始时玩家A拿走1，则排列变为[2,100,4]，接下来玩家B可以拿走2或4，然后继续轮到玩家A...
//如果开始时玩家A拿走4，则排列变为[1,2,100]，接下来玩家B可以拿走1或100，然后继续轮到玩家A...
//玩家A作为绝顶聪明的人不会先拿4，因为拿4之后，玩家B将拿走100。所以玩家A会先拿1，让排列变为[2,100,4]，
//接下来玩家B不管怎么选，100都会被玩家A拿走。玩家A会获胜，分数为101。所以返回101。
//arr=[1,100,2]。
//开始时，玩家A不管拿1还是2，玩家B作为绝顶聪明的人，都会把100拿走。玩家B会获胜，分数为100。所以返回100。

func findMaxForCardInLine(input []int) int {
	if f(input, 0, len(input)-1) > s(input, 0, len(input)-1) {
		return f(input, 0, len(input)-1)
	} else {
		return s(input, 0, len(input)-1)
	}
}

// 先手函数
func f(input []int, L, R int) int {
	if L == R {
		return input[L] // baseCase: 如果只有一个元素，直接返回
	}
	// 由于只能取左右两侧元素，所以比较 当取最左侧元素后 + [L+1, R]的后手函数 和 当取最右侧元素后 + [L, R-1]的后手函数
	// 	哪个大取哪个。
	if input[L]+s(input, L+1, R) > input[R]+s(input, L, R-1) {
		return input[L] + s(input, L+1, R)
	} else {
		return input[R] + s(input, L, R-1)
	}
}

// 后手函数
func s(input []int, L, R int) int {
	if L == R {
		return 0 // baseCase: 如果只有一个元素，返回0，因为元素被先手拿了
	}
	// 由于后手函数取哪个取决于先手函数，所以后手函数取的是较小的那个
	if f(input, L+1, R) > f(input, L, R-1) {
		return f(input, L, R-1)
	} else {
		return f(input, L+1, R)
	}
}
