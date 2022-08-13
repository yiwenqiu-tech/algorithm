package sort

import (
	"reflect"
	"testing"
)

func TestRadixSort(t *testing.T) {
	type args struct {
		array []int
		radix int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test radix sort",
			args: args{
				array: []int{9, 5, 2, 66, 2, 101},
				radix: 10,
			},
			want: []int{2, 2, 5, 9, 66, 101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RadixSort(tt.args.array, tt.args.radix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RadixSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
