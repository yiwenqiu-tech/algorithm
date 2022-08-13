package sort

import (
	"reflect"
	"testing"
)

func TestSortArrayForAlmostSortedArray(t *testing.T) {
	type args struct {
		array []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test sort array for almost sorted array",
			args: args{
				array: []int{4, 2, 1, 5, 7},
				k:     5,
			},
			want: []int{1, 2, 4, 5, 7},
		},
		{
			name: "test sort array for almost sorted array",
			args: args{
				array: []int{3, 2, 1, 6, 5, 4},
				k:     3,
			},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortArrayForAlmostSortedArray(tt.args.array, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrayForAlmostSortedArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
