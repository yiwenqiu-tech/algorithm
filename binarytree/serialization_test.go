package binarytree

import (
	"reflect"
	"testing"
)

func TestSerialization(t *testing.T) {
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test serial",
			args: args{
				head: &Node{
					Value: 1,
					Left: &Node{
						Value: 2,
						Left: &Node{
							Value: 4,
						},
						Right: &Node{
							Value: 5,
						},
					},
					Right: &Node{
						Value: 3,
						Left: &Node{
							Value: 6,
						},
						Right: &Node{
							Value: 7,
						},
					},
				},
			},
			want: "1_2_4_#_#_5_#_#_3_6_#_#_7_#_#_",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Serialization(tt.args.head); got != tt.want {
				t.Errorf("Serialization() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeSerialization(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			name: "test deSerialization",
			args: args{
				str: "1_2_4_#_#_5_#_#_3_6_#_#_7_#_#_",
			},
			want: &Node{
				Value: 1,
				Left: &Node{
					Value: 2,
					Left: &Node{
						Value: 4,
					},
					Right: &Node{
						Value: 5,
					},
				},
				Right: &Node{
					Value: 3,
					Left: &Node{
						Value: 6,
					},
					Right: &Node{
						Value: 7,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DeSerialization(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeSerialization() = %v, want %v", got, tt.want)
			}
		})
	}
}
