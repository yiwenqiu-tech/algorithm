package backtracking

import "testing"

func Test_convertToLetterString(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test convert to letter string",
			args: args{
				input: "111",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToLetterString(tt.args.input); got != tt.want {
				t.Errorf("convertToLetterString() = %v, want %v", got, tt.want)
			}
		})
	}
}
