package xor

import "testing"

func TestFindOnlyOneOddNumber(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find only one odd number",
			args: args{
				[]int{1, 1, 1, 2, 2, 3, 4, 4, 5, 5, 3},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindOnlyOneOddNumber(tt.args.array); got != tt.want {
				t.Errorf("FindOnlyOneOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindOnlyTwoOddNumber(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name: "find only two odd number",
			args: args{
				array: []int{1, 1, 2, 2, 3, 4},
			},
			want:  3,
			want1: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := FindOnlyTwoOddNumber(tt.args.array)
			if got != tt.want {
				t.Errorf("FindOnlyTwoOddNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FindOnlyTwoOddNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
