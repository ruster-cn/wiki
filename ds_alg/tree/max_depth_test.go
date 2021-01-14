package tree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "unit test",
			args: args{root: &Node{
				Val:   3,
				Right: &Node{Val: 9},
				Left:  &Node{Val: 20, Right: &Node{Val: 15}, Left: &Node{Val: 7}},
			}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
