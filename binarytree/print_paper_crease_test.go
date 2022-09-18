package binarytree

import "testing"

func TestPrintPaperCrease(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "print paper crease",
			args: args{n: 1},
			want: "down ",
		},
		{
			name: "print paper crease",
			args: args{n: 2},
			want: "down down up ",
		},
		{
			name: "print paper crease",
			args: args{n: 3},
			want: "down down up down down up up ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrintPaperCrease(tt.args.n); got != tt.want {
				t.Errorf("PrintPaperCrease() = %v, want %v", got, tt.want)
			}
		})
	}
}
