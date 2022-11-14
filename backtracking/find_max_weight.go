package backtracking

import "fmt"

// 给定两个长度都为N的数组weights和values，weights[i]和values[i]分别代表i号物品的重量和价值。
//	给定一个正数bag，表示一个载重bag的袋子，你装的物品不能超过这个重量。返回你能装下最多的价值是多少？
// 	解体思路：一个一个试，每个有要和不要的选择

func FindMaxWeight(weights []int, values []int, bag int) int {
	return processForFindMaxWeight(weights, values, 0,
		0, 0, bag, []int{})
}

func processForFindMaxWeight(weights []int, values []int, i int,
	alreadyWeight int, alreadyValues int, bag int, chooseI []int) int {
	if alreadyWeight > bag { // 超过包的限制了
		return 0
	}
	if i == len(weights) { // 选择完了，直接返回alreadyValues
		fmt.Printf("alreadyValues: %v, choose: %+v\n", alreadyValues, chooseI)
		return alreadyValues
	}

	copyArrayChoose := copyArray(chooseI)
	copyArrayChoose = append(copyArrayChoose, i)
	chooseIValues := processForFindMaxWeight(weights, values, i+1,
		alreadyWeight+weights[i], alreadyValues+values[i], bag, copyArrayChoose)

	copyArrayNoChoose := copyArray(chooseI)
	noChooseIValues := processForFindMaxWeight(weights, values, i+1,
		alreadyWeight, alreadyValues, bag, copyArrayNoChoose)

	if chooseIValues > noChooseIValues {
		return chooseIValues
	} else {
		return noChooseIValues
	}
}

func copyArray(chooseI []int) []int {
	var res []int
	for i := range chooseI {
		res = append(res, chooseI[i])
	}
	return res
}
