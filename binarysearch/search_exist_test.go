package binarysearch

import "testing"

func TestSearchExist(t *testing.T) {
	type args struct {
		array []int
		num   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "have no num",
			args: args{
				array: []int{},
				num:   0,
			},
			want: false,
		},
		{
			name: "have odd num",
			args: args{
				array: []int{1, 3, 5, 7, 9},
				num:   7,
			},
			want: true,
		},
		{
			name: "have odd num, but not exist",
			args: args{
				array: []int{1, 3, 5, 7, 9},
				num:   10,
			},
			want: false,
		},
		{
			name: "have even num",
			args: args{
				array: []int{1, 3, 5, 7, 9, 11},
				num:   1,
			},
			want: true,
		},
		{
			name: "have even num, but not exist",
			args: args{
				array: []int{1, 3, 5, 7, 9, 11},
				num:   12,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchExist(tt.args.array, tt.args.num); got != tt.want {
				t.Errorf("SearchExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
