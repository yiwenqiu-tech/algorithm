package trietree

import (
	"testing"
)

func Test_node_Insert(t *testing.T) {
	type fields struct {
		Pass int
		End  int
		Next []*node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test insert abc",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &node{
				Pass: tt.fields.Pass,
				End:  tt.fields.End,
				Next: tt.fields.Next,
			}
			root.Insert(tt.args.word)
			if root.Pass != 1 {
				t.Errorf("root pass() = %v, want %v", root.Pass, 1)
			}
			if root.Next[0] == nil {
				t.Errorf("root Next[0] nil, want not nil")
			}
			if root.Next[1] != nil {
				t.Errorf("root Next[0] not nil, want nil")
			}
			if root.Next[0].Pass != 1 {
				t.Errorf("root Next[0] pass() = %v, want %v", root.Next[0].Pass, 1)
			}
			if root.Next[0].Next[1] == nil {
				t.Errorf("root.Next[0].Next[1] nil, want not nil")
			}
			if root.Next[0].Next[1].Pass != 1 {
				t.Errorf("root.Next[0].Next[1] pass() = %v, want %v", root.Next[0].Next[1].Pass, 1)
			}
			if root.Next[0].Next[1].Next[2] == nil {
				t.Errorf("root.Next[0].Next[1].Next[2] nil, want not nil")
			}
			if root.Next[0].Next[1].Next[2].Pass != 1 {
				t.Errorf("root.Next[0].Next[1].Next[2] pass() = %v, want %v", root.Next[0].Next[1].Next[2].Pass, 1)
			}
			if root.Next[0].Next[1].Next[2].End != 1 {
				t.Errorf("root.Next[0].Next[1].Next[2] end() = %v, want %v", root.Next[0].Next[1].Next[2].End, 1)
			}
		})
	}
}

func Test_node_Search(t *testing.T) {
	type fields struct {
		Pass int
		End  int
		Next []*node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test search",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "abc",
			},
			want: 2,
		},
		{
			name: "test search",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "abd",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &node{
				Pass: tt.fields.Pass,
				End:  tt.fields.End,
				Next: tt.fields.Next,
			}
			root.Insert("abc")
			root.Insert("abc")
			if got := root.Search(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_PrefixMatch(t *testing.T) {
	type fields struct {
		Pass int
		End  int
		Next []*node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "test prefix",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "abc",
			},
			want: 2,
		},
		{
			name: "test prefix",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "abd",
			},
			want: 0,
		},
		{
			name: "test prefix",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "ab",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &node{
				Pass: tt.fields.Pass,
				End:  tt.fields.End,
				Next: tt.fields.Next,
			}
			root.Insert("abc")
			root.Insert("abcdae")
			root.Insert("ab")
			if got := root.PrefixMatch(tt.args.word); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_Delete(t *testing.T) {
	type fields struct {
		Pass int
		End  int
		Next []*node
	}
	type args struct {
		word string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "test delete",
			fields: fields{
				Pass: 0,
				End:  0,
				Next: make([]*node, 26),
			},
			args: args{
				word: "abc",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := &node{
				Pass: tt.fields.Pass,
				End:  tt.fields.End,
				Next: tt.fields.Next,
			}
			root.Insert("abc")
			exist := root.IsExist("abc")
			if !exist {
				t.Errorf("isSearch = %v, want %v", exist, true)
			}
			root.Delete("abc")
			exist = root.IsExist("abc")
			if exist {
				t.Errorf("isSearch = %v, want %v", exist, false)
			}
		})
	}
}
