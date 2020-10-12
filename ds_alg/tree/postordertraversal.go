package tree

func postorderTraversalLoop(root *Node, list *[]int) {
	if root == nil {
		return
	}
	var last *Node
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		//检查弹出的node无右孩子或者右孩子为
		if node.Right == nil || node.Right == last {
			stack = stack[:len(stack)-1]
			*list = append(*list, node.Val)
			last = node
		} else {
			root = node.Right
		}

	}
}
