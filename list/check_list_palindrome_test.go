package list

import "testing"

func TestCheckListPalindromeByStack(t *testing.T) {
	type args struct {
		head *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check list palindrome by stack",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 1,
							Next:  nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 2,
							Next: &SingleLinkedNode{
								Value: 1,
								Next:  nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 15,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 15,
								Next: &SingleLinkedNode{
									Value: 1,
									Next:  nil,
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 15,
								Next: &SingleLinkedNode{
									Value: 1,
									Next:  nil,
								},
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckListPalindromeByStack(tt.args.head); got != tt.want {
				t.Errorf("CheckListPalindromeByStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckListPalindromeByStack2(t *testing.T) {
	type args struct {
		head *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check list palindrome by stack2",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack2",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 1,
							Next:  nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack2",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 2,
							Next: &SingleLinkedNode{
								Value: 1,
								Next:  nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack2",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 15,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 15,
								Next: &SingleLinkedNode{
									Value: 1,
									Next:  nil,
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "check list palindrome by stack2",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 15,
								Next: &SingleLinkedNode{
									Value: 1,
									Next:  nil,
								},
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckListPalindromeByStack2(tt.args.head); got != tt.want {
				t.Errorf("CheckListPalindromeByStack2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckListPalindrome(t *testing.T) {
	type args struct {
		head *SingleLinkedNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "CheckListPalindrome",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next:  nil,
				},
			},
			want: true,
		},
		{
			name: "CheckListPalindrome",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 1,
							Next:  nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "CheckListPalindrome",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 2,
							Next: &SingleLinkedNode{
								Value: 1,
								Next:  nil,
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "CheckListPalindrome",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 15,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 15,
								Next: &SingleLinkedNode{
									Value: 1,
									Next:  nil,
								},
							},
						},
					},
				},
			},
			want: true,
		},
		{
			name: "CheckListPalindrome",
			args: args{
				head: &SingleLinkedNode{
					Value: 1,
					Next: &SingleLinkedNode{
						Value: 2,
						Next: &SingleLinkedNode{
							Value: 10,
							Next: &SingleLinkedNode{
								Value: 15,
								Next: &SingleLinkedNode{
									Value: 1,
									Next:  nil,
								},
							},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckListPalindrome(tt.args.head); got != tt.want {
				t.Errorf("CheckListPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
