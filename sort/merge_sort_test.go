package sort

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test Merge sort",
			args: args{[]int{3, 4, 6, 1, 2}},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "test Merge sort",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "test Merge sort",
			args: args{[]int{2, 1}},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSort2(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test Merge sort2",
			args: args{[]int{3, 4, 6, 1, 2, 9}},
			want: []int{1, 2, 3, 4, 6, 9},
		},
		{
			name: "test Merge sort2",
			args: args{[]int{3, 4, 6, 1, 2}},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "test Merge sort2",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "test Merge sort2",
			args: args{[]int{2, 1}},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort2(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
