package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test Bubble sort",
			args: args{[]int{3, 4, 6, 1, 2}},
			want: []int{1, 2, 3, 4, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 对数器实现
// 	 1.实现一个相对简单（比较容易想到的）的方法b，这里以golang内置的排序作为方法b
//	 2.实现一个样例生成器（随机、大量的生成）
//   3.运行方法a（待测）与方法b，并对每组结果作为对比
func TestBubbleSortChecker(t *testing.T) {
	var testCase [][]int
	for i := 0; i < 100; i++ {
		testCase = append(testCase, genRandIntArr(10))
	}

	for _, tt := range testCase {
		if got := BubbleSort(tt); !reflect.DeepEqual(got, golangSort(tt)) {
			t.Errorf("BubbleSort() = %v, want %v", got, golangSort(tt))
		}
	}
}

// 通过传入的长度，生成100内的int类型随机数组
func genRandIntArr(length int) []int {
	nums := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(100)
	}
	return nums
}
