package backtracking

import (
	"reflect"
	"testing"
)

func Test_printAllPermutations(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "print all permutations",
			args: args{
				input: "abc",
			},
			want: []string{
				"abc",
				"acb",
				"bac",
				"bca",
				"cba",
				"cab",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printAllPermutations(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printAllPermutations() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printAllPermutationsWithoutDup(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test printAllPermutationsWithoutDup",
			args: args{
				input: "bbb",
			},
			want: []string{"bbb"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printAllPermutationsWithoutDup(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printAllPermutationsWithoutDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
