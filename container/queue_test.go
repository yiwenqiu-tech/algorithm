package container

import (
	"reflect"
	"testing"
)

func TestQueueBySlice(t *testing.T) {
	stack := &QueueBySlice{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(9)
	stack.Push(8)
	if !reflect.DeepEqual(stack.values, []interface{}{1, 2, 3, 9, 8}) {
		t.Errorf("push result = %v, want %v", stack.values, []interface{}{1, 2, 3, 9, 8})
	}
	if !reflect.DeepEqual(stack.Len(), 5) {
		t.Errorf("len result = %v, want %v", stack.Len(), 5)
	}
	res := stack.Pop()
	if !reflect.DeepEqual(res, 1) {
		t.Errorf("pop result = %v, want %v", res, 1)
	}
	if !reflect.DeepEqual(stack.Len(), 4) {
		t.Errorf("len result = %v, want %v", stack.Len(), 4)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 2) {
		t.Errorf("pop result = %v, want %v", res, 2)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 3) {
		t.Errorf("pop result = %v, want %v", res, 3)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 9) {
		t.Errorf("pop result = %v, want %v", res, 9)
	}
	if !reflect.DeepEqual(stack.Empty(), false) {
		t.Errorf("empty result = %v, want %v", stack.Empty(), false)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 8) {
		t.Errorf("pop result = %v, want %v", res, 8)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, nil) {
		t.Errorf("pop result = %v, want %v", res, nil)
	}
	if !reflect.DeepEqual(stack.Empty(), true) {
		t.Errorf("empty result = %v, want %v", stack.Empty(), true)
	}
}

func TestQueueByList(t *testing.T) {
	stack := &QueueByList{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(9)
	stack.Push(8)
	if !reflect.DeepEqual(stack.Len(), 5) {
		t.Errorf("len result = %v, want %v", stack.Len(), 5)
	}
	res := stack.Pop()
	if !reflect.DeepEqual(res, 1) {
		t.Errorf("pop result = %v, want %v", res, 1)
	}
	if !reflect.DeepEqual(stack.Len(), 4) {
		t.Errorf("len result = %v, want %v", stack.Len(), 4)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 2) {
		t.Errorf("pop result = %v, want %v", res, 2)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 3) {
		t.Errorf("pop result = %v, want %v", res, 3)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 9) {
		t.Errorf("pop result = %v, want %v", res, 9)
	}
	if !reflect.DeepEqual(stack.Empty(), false) {
		t.Errorf("empty result = %v, want %v", stack.Empty(), false)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, 8) {
		t.Errorf("pop result = %v, want %v", res, 8)
	}
	res = stack.Pop()
	if !reflect.DeepEqual(res, nil) {
		t.Errorf("pop result = %v, want %v", res, nil)
	}
	if !reflect.DeepEqual(stack.Empty(), true) {
		t.Errorf("empty result = %v, want %v", stack.Empty(), true)
	}
}
