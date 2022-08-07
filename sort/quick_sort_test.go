package sort

import (
	"reflect"
	"testing"
)

func TestQuickSortV1(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "quick sort V1 test",
			args: args{
				array: []int{5, 4, 8, 7, 1},
			},
			want: []int{1, 4, 5, 7, 8},
		},
		{
			name: "quick sort V1 test",
			args: args{
				array: []int{5},
			},
			want: []int{5},
		},
		{
			name: "quick sort V1 test",
			args: args{
				array: []int{5, 4, 8, 7, 1, 9},
			},
			want: []int{1, 4, 5, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortV1(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortV1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortV2(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "quick sort V2 test",
			args: args{
				array: []int{5, 4, 8, 7, 1},
			},
			want: []int{1, 4, 5, 7, 8},
		},
		{
			name: "quick sort V2 test",
			args: args{
				array: []int{5},
			},
			want: []int{5},
		},
		{
			name: "quick sort V2 test",
			args: args{
				array: []int{5, 4, 8, 7, 1, 9},
			},
			want: []int{1, 4, 5, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortV2(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortV2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQuickSortV3(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "quick sort V3 test",
			args: args{
				array: []int{5, 4, 8, 7, 1},
			},
			want: []int{1, 4, 5, 7, 8},
		},
		{
			name: "quick sort V3 test",
			args: args{
				array: []int{5},
			},
			want: []int{5},
		},
		{
			name: "quick sort V3 test",
			args: args{
				array: []int{5, 4, 8, 7, 1, 9},
			},
			want: []int{1, 4, 5, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSortV3(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSortV3() = %v, want %v", got, tt.want)
			}
		})
	}
}
