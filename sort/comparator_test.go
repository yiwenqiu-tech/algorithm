package sort

import (
	"reflect"
	"testing"
)

func TestCompareStudent1(t *testing.T) {
	type args struct {
		studentSlice []Student
	}
	tests := []struct {
		name string
		args args
		want []Student
	}{
		{
			name: "compare students",
			args: args{
				studentSlice: []Student{
					{
						Age:  28,
						Name: "xiaowang",
					},
					{
						Age:  27,
						Name: "xiaocai",
					},
					{
						Age:  29,
						Name: "xiaoming",
					},
				},
			},
			want: []Student{
				{
					Age:  27,
					Name: "xiaocai",
				},
				{
					Age:  28,
					Name: "xiaowang",
				},
				{
					Age:  29,
					Name: "xiaoming",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareStudent(tt.args.studentSlice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareStudent() = %v, want %v", got, tt.want)
			}
		})
	}
}
