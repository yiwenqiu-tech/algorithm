package sort

import (
	"reflect"
	"testing"
)

func TestDutchNationalFlagSimplifiedVersion(t *testing.T) {
	type args struct {
		array []int
		num   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test DutchNationalFlagSimplifiedVersion when only a num",
			args: args{
				array: []int{1},
				num:   4,
			},
			want: []int{1},
		},
		{
			name: "test DutchNationalFlagSimplifiedVersion",
			args: args{
				array: []int{5, 2, 6, 1, 3},
				num:   4,
			},
			want: []int{2, 1, 3, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DutchNationalFlagSimplifiedVersion(tt.args.array, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DutchNationalFlagSimplifiedVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDutchNationalFlag(t *testing.T) {
	type args struct {
		array []int
		num   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test DutchNationalFlag when only a num",
			args: args{
				array: []int{1},
				num:   4,
			},
			want: []int{1},
		},
		{
			name: "test DutchNationalFlag",
			args: args{
				array: []int{3, 9, 2, 4, 6, 7},
				num:   4,
			},
			want: []int{3, 2, 4, 6, 7, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DutchNationalFlag(tt.args.array, tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DutchNationalFlag() = %v, want %v", got, tt.want)
			}
		})
	}
}
