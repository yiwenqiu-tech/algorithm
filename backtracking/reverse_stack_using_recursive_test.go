package backtracking

import (
	"algorithm/container"
	"testing"
)

func Test_getStackBottom(t *testing.T) {
	stack := &container.StackBySlice{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	type args struct {
		stack *container.StackBySlice
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test getStackBottom",
			args: args{
				stack: stack,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getStackBottom(tt.args.stack); got != tt.want {
				t.Errorf("getStackBottom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseStackUsingRecursive(t *testing.T) {
	stack := &container.StackBySlice{}
	stack.Push(3)
	stack.Push(2)
	stack.Push(1)

	ReverseStackUsingRecursive(stack)

	top := stack.Pop().(int)
	if top != 3 {
		t.Errorf("reverse top() = %v, want %v", top, 3)
	}

	top = stack.Pop().(int)
	if top != 2 {
		t.Errorf("reverse second() = %v, want %v", top, 2)
	}

	top = stack.Pop().(int)
	if top != 1 {
		t.Errorf("reverse third() = %v, want %v", top, 1)
	}

}
