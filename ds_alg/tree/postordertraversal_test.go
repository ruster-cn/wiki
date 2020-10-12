package tree

import "testing"

func Test_postorderTraversalLoop(t *testing.T) {
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
			postorderTraversalLoop(tt.args.root, &tt.args.list)
			t.Logf("%#v", tt.args.list)
		})
	}
}
