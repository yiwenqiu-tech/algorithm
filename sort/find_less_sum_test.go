package sort

import "testing"

func TestFindLessSum(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find less sum",
			args: args{
				array: []int{1, 3, 4, 2, 5},
			},
			want: 16,
		},
		{
			name: "find less sum",
			args: args{
				array: []int{1},
			},
			want: 0,
		},
		{
			name: "find less sum",
			args: args{
				array: []int{2, 1},
			},
			want: 0,
		},
		{
			name: "find less sum",
			args: args{
				array: []int{1, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindLessSum(tt.args.array); got != tt.want {
				t.Errorf("FindLessSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
