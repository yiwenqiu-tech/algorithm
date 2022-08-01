package sort

import "testing"

func Test_swapByTemp(t *testing.T) {
	type args struct {
		array []int
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test swap by temp",
			args: args{
				array: []int{3, 10, 1},
				i:     0,
				j:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapByTemp(tt.args.array, tt.args.i, tt.args.j)
			if tt.args.array[0] != 1 {
				t.Errorf("array[0] = %v, want %v", tt.args.array[0], 1)
			}
			if tt.args.array[2] != 3 {
				t.Errorf("array[0] = %v, want %v", tt.args.array[2], 3)
			}
		})
	}
}

func Test_swapByXOR(t *testing.T) {
	type args struct {
		array []int
		i     int
		j     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test swap by xor",
			args: args{
				array: []int{3, 10, 1},
				i:     0,
				j:     2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			swapByXOR(tt.args.array, tt.args.i, tt.args.j)
			if tt.args.array[0] != 1 {
				t.Errorf("array[0] = %v, want %v", tt.args.array[0], 1)
			}
			if tt.args.array[2] != 3 {
				t.Errorf("array[0] = %v, want %v", tt.args.array[2], 3)
			}
		})
	}
}
