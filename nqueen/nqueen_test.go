package nqueen

import "testing"

func TestNQueen(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test n queen",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "test n queen",
			args: args{n: 2},
			want: 0,
		},
		{
			name: "test n queen",
			args: args{n: 8},
			want: 92,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NQueen(tt.args.n); got != tt.want {
				t.Errorf("NQueen() = %v, want %v", got, tt.want)
			}
		})
	}
}
