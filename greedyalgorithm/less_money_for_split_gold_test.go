package greedyalgorithm

import "testing"

func TestLessMoneyForSplitGoldByExhaustive(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test less money for split gold",
			args: args{
				[]int{10, 20, 30},
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LessMoneyForSplitGoldByExhaustive(tt.args.input); got != tt.want {
				t.Errorf("LessMoneyForSplitGoldByExhaustive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLessMoneyForSplitGold(t *testing.T) {
	maxSize := 6
	maxItem := 1000000
	loopTime := 10000
	for i := 0; i < loopTime; i++ {
		input := generateInput(maxSize, maxItem)
		if LessMoneyForSplitGoldByExhaustive(input) != LessMoneyForSplitGold(input) {
			t.Errorf("LessMoneyForSplitGoldByExhaustive() = %v, LessMoneyForSplitGold %v",
				LessMoneyForSplitGoldByExhaustive(input), LessMoneyForSplitGold(input))
		}
	}
}

func TestLessMoneyForSplitGold1(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test less money for split gold",
			args: args{
				[]int{10, 20, 30},
			},
			want: 90,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LessMoneyForSplitGold(tt.args.input); got != tt.want {
				t.Errorf("LessMoneyForSplitGold() = %v, want %v", got, tt.want)
			}
		})
	}
}
