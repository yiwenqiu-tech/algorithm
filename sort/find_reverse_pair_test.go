package sort

import "testing"

func TestFindReversePair(t *testing.T) {
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
			want: 2,
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
			want: 1,
		},
		{
			name: "find less sum",
			args: args{
				array: []int{1, 2},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindReversePair(tt.args.array); got != tt.want {
				t.Errorf("FindReversePair() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckFindReversePair(t *testing.T) {
	// 生成随机输入
	var inputs [][]int
	for i := 0; i < 100; i++ {
		inputs = append(inputs, genRandIntArr(100))
	}

	for _, tt := range inputs {
		checker := FindReversePair2(tt)
		if got := FindReversePair(tt); got != checker {
			t.Errorf("FindReversePair2() = %v, want %v", got, checker)
		}
	}
}
