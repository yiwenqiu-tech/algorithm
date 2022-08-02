package binarysearch

import "testing"

func TestSearchNearRight(t *testing.T) {
	type args struct {
		array []int
		num   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "have no num",
			args: args{
				array: []int{},
				num:   0,
			},
			want: -1,
		},
		{
			name: "have odd num, and exist num",
			args: args{
				array: []int{1, 3, 5, 7, 9},
				num:   7,
			},
			want: 3,
		},
		{
			name: "have odd num, but not more than num",
			args: args{
				array: []int{1, 3, 5, 7, 9},
				num:   10,
			},
			want: 4,
		},
		{
			name: "have odd num, but not exist num",
			args: args{
				array: []int{1, 3, 5, 7, 9},
				num:   6,
			},
			want: 2,
		},
		{
			name: "have more than 1 num",
			args: args{
				array: []int{1, 3, 3, 3, 3, 11},
				num:   3,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SearchNearRight(tt.args.array, tt.args.num); got != tt.want {
				t.Errorf("SearchNearRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
