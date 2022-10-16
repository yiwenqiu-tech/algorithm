package greedyalgorithm

import "testing"

func TestIPO(t *testing.T) {
	type args struct {
		costs   []int
		profits []int
		k       int
		m       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test ipo",
			args: args{
				costs:   []int{1},
				profits: []int{1},
				k:       2,
				m:       0,
			},
			want: 0,
		},
		{
			name: "test ipo",
			args: args{
				costs:   []int{1},
				profits: []int{1},
				k:       2,
				m:       1,
			},
			want: 2,
		},
		{
			name: "test ipo",
			args: args{
				costs:   []int{1, 1},
				profits: []int{1, 4},
				k:       1,
				m:       1,
			},
			want: 5,
		},
		{
			name: "test ipo",
			args: args{
				costs:   []int{1, 1, 2, 4},
				profits: []int{1, 4, 3, 6},
				k:       3,
				m:       1,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPO(tt.args.costs, tt.args.profits, tt.args.k, tt.args.m); got != tt.want {
				t.Errorf("IPO() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIPOByExhaustive(t *testing.T) {
	type args struct {
		costs   []int
		profits []int
		k       int
		m       int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//{
		//	name: "test ipo",
		//	args: args{
		//		costs: []int{1},
		//		profits: []int{1},
		//		k: 2,
		//		m:0,
		//	},
		//	want: 0,
		//},
		//{
		//	name: "test ipo",
		//	args: args{
		//		costs: []int{1},
		//		profits: []int{1},
		//		k: 2,
		//		m: 1,
		//	},
		//	want: 2,
		//},
		//{
		//	name: "test ipo",
		//	args: args{
		//		costs: []int{1, 1},
		//		profits: []int{1, 4},
		//		k: 1,
		//		m: 1,
		//	},
		//	want: 5,
		//},
		{
			name: "test ipo",
			args: args{
				costs:   []int{1, 1, 2, 4},
				profits: []int{1, 4, 3, 6},
				k:       3,
				m:       1,
			},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IPOByExhaustive(tt.args.costs, tt.args.profits, tt.args.k, tt.args.m); got != tt.want {
				t.Errorf("IPOByExhaustive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLessMoneyForIPO(t *testing.T) {
	maxSize := 8
	maxItem := 1000000
	loopTime := 1000
	for i := 0; i < loopTime; i++ {
		costs, profits := generateInputIPO(maxSize, maxItem)
		k := generateK(maxSize)
		m := costs[0]
		if IPO(costs, profits, k, m) != IPOByExhaustive(costs, profits, k, m) {
			t.Errorf("IPO() = %v, IPOByExhaustive %v",
				IPO(costs, profits, k, m), IPOByExhaustive(costs, profits, k, m))
		}
	}
}
