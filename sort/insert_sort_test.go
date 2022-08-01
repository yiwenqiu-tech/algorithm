package sort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test Insert sort",
			args: args{[]int{3, 4, 6, 1, 2}},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "test Insert sort",
			args: args{[]int{1}},
			want: []int{1},
		},
		{
			name: "test Insert sort",
			args: args{[]int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
