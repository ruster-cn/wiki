package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPreOrderTraversalRecursion(t *testing.T) {
	type args struct {
		root *Node
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test",
			args: args{
				root: &Node{
					Left: &Node{
						Left: &Node{
							Val: 4,
						},
						Val: 2,
						Right: &Node{
							Val: 5,
						},
					},
					Val: 1,
					Right: &Node{
						Left: &Node{
							Val: 6,
						},
						Val: 3,
						Right: &Node{
							Val: 7,
						},
					},
				},
				list: make([]int, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PreOrderTraversalRecursion(tt.args.root, &tt.args.list)
			t.Logf("list is: %#v", tt.args.list)
		})
	}
}

func TestPreOrderTraversalLoop(t *testing.T) {
	type args struct {
		root *Node
		list []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test",
			args: args{
				root: &Node{
					Left: &Node{
						Left: &Node{
							Val: 4,
						},
						Val: 2,
						Right: &Node{
							Val: 5,
						},
					},
					Val: 1,
					Right: &Node{
						Left: &Node{
							Val: 6,
						},
						Val: 3,
						Right: &Node{
							Val: 7,
						},
					},
				},
				list: make([]int, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Printf("%p \n", tt.args.list)
			PreOrderTraversalLoop(tt.args.root, &tt.args.list)
			t.Logf("list is: %#v", tt.args.list)
		})
	}
}

func TestPreOrderTraversalDivide(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test",
			args: args{
				root: &Node{
					Left: &Node{
						Left: &Node{
							Val: 4,
						},
						Val: 2,
						Right: &Node{
							Val: 5,
						},
					},
					Val: 1,
					Right: &Node{
						Left: &Node{
							Val: 6,
						},
						Val: 3,
						Right: &Node{
							Val: 7,
						},
					},
				},
			},
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PreOrderTraversalDivide(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreOrderTraversalDivide() = %v, want %v", got, tt.want)
			}
		})
	}
}
