package tree

func inorderTraversalLoop(root *Node, list *[]int) {
	if root == nil {
		return
	}
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			*list = append(*list, node.Val)
			root = node.Right
		}
	}
}
