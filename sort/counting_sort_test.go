package sort

import (
	"reflect"
	"testing"
)

func TestCountSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test Count sort",
			args: args{[]int{3, 4, 6, 1, 2}},
			want: []int{1, 2, 3, 4, 6},
		},
		{
			name: "test Count sort",
			args: args{[]int{3}},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
