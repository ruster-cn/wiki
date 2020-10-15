package tree

func levelOrder(root *Node, list *[]int) {
	if root == nil {
		return
	}
	first := make([]*Node, 0)
	second := make([]*Node, 0)
	first = append(first, root)
	for len(first) != 0 || len(second) != 0 {
		for len(first) != 0 {
			node := first[0]
			first = first[1:]
			*list = append(*list, node.Val)
			if node.Left != nil {
				second = append(second, node.Left)
			}
			if node.Right != nil {
				second = append(second, node.Right)
			}
		}
		for len(second) != 0 {
			node := second[0]
			second = second[1:]
			*list = append(*list, node.Val)
			if node.Left != nil {
				first = append(first, node.Left)
			}
			if node.Right != nil {
				first = append(first, node.Right)
			}
		}
	}

}
