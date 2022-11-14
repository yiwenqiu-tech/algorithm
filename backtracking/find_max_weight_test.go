package backtracking

import "testing"

func TestFindMaxWeight(t *testing.T) {
	type args struct {
		weights []int
		values  []int
		bag     int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test max weight",
			args: args{
				weights: []int{3, 2, 4, 7},
				values:  []int{5, 6, 3, 19},
				bag:     11,
			},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMaxWeight(tt.args.weights, tt.args.values, tt.args.bag); got != tt.want {
				t.Errorf("FindMaxWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
