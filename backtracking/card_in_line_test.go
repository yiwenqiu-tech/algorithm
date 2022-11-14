package backtracking

import "testing"

func Test_findMaxForCardInLine(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test find max for card in line",
			args: args{
				input: []int{1, 100, 2},
			},
			want: 100,
		},
		{
			name: "test find max for card in line",
			args: args{
				input: []int{1, 2, 100, 4},
			},
			want: 101,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxForCardInLine(tt.args.input); got != tt.want {
				t.Errorf("findMaxForCardInLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
