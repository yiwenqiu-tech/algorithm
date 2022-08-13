package sort

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func Test_heapInsert(t *testing.T) {
	type args struct {
		array []int
		index int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test heapInsert",
			args: args{
				array: []int{1},
				index: 0,
			},
			want: []int{1},
		},
		{
			name: "test heapInsert",
			args: args{
				array: []int{1},
				index: 1,
			},
			want: []int{1},
		},
		{
			name: "test heapInsert",
			args: args{
				array: []int{1, 2},
				index: 1,
			},
			want: []int{2, 1},
		},
		{
			name: "test heapInsert",
			args: args{
				array: []int{2, 1, 3},
				index: 2,
			},
			want: []int{3, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapInsert(tt.args.array, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_heapify(t *testing.T) {
	type args struct {
		array    []int
		index    int
		heapSize int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test heapify, invalid input",
			args: args{
				array:    []int{1},
				index:    1,
				heapSize: 1,
			},
			want: []int{1},
		},
		{
			name: "test heapify, invalid input",
			args: args{
				array:    []int{2, 1},
				index:    1,
				heapSize: 3,
			},
			want: []int{2, 1},
		},
		{
			name: "test heapify",
			args: args{
				array:    []int{2, 1, 3},
				index:    0,
				heapSize: 3,
			},
			want: []int{3, 1, 2},
		},
		{
			name: "test heapify",
			args: args{
				array:    []int{1, 6, 5, 4, 3, 2, 7},
				index:    0,
				heapSize: 6,
			},
			want: []int{6, 4, 5, 1, 3, 2, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := heapify(tt.args.array, tt.args.index, tt.args.heapSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("heapify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test heap sort",
			args: args{
				array: []int{1},
			},
			want: []int{1},
		},
		{
			name: "test heap sort",
			args: args{
				array: []int{3, 2, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "test heap sort",
			args: args{
				array: []int{4, 8, 1, 2},
			},
			want: []int{1, 2, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HeapSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckHeapSort(t *testing.T) {
	// 生成随机输入
	var inputs1 = make([][]int, 100)
	var inputs2 = make([][]int, 100)
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			rand.Seed(time.Now().UnixNano())
			randNum := rand.Intn(100)
			inputs1[i] = append(inputs1[i], randNum)
			inputs2[i] = append(inputs2[i], randNum)
		}
	}

	for index := range inputs1 {
		checker := GoHeapSort(inputs1[index])
		got := HeapSort(inputs2[index])
		if !reflect.DeepEqual(got, checker) {
			t.Errorf("HeapSort() = %v, want %v", got, checker)
		}
	}
}
