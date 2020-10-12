package tree

//NOTE:二叉树前序遍历，先访问根节点再访问左子树，再访问右子树。左右子树遍历时依然按照上述规则。

//PreOrderTraversalRecursion 递归版本前序遍历
func PreOrderTraversalRecursion(root *Node, list *[]int) {
	if root == nil {
		return
	}
	//NOTE: 这个[]int必须使用指针，是因为append如果超出了原数组长度，需要重新分配内存
	//这时候最上层递归函数和下层递归函数的list指向的内存块就不一致了。
	*list = append(*list, root.Val)
	PreOrderTraversalRecursion(root.Left, list)
	PreOrderTraversalRecursion(root.Right, list)
}

//PreOrderTraversalLoop 非递归版本前序遍历
func PreOrderTraversalLoop(root *Node, list *[]int) {
	if root == nil {
		return
	}
	stack := make([]*Node, 0)
	for root != nil || len(stack) != 0 {
		for root != nil {
			//NOTE: append 方法会改变slice的内存地址，方法使用slice 参数做结果收集时一定要使用，slice指针
			*list = append(*list, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
}
