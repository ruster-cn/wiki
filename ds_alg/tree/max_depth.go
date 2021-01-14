package tree

/*
	## 思路
		右子树高度 = maxDepth(root.Right) + 1
		左子树高度 = maxDepth(root.Left) + 1
		返回左子树和右子树最大值
	link: https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/
*/
func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	right, left := 1, 1

	if root.Left != nil {
		left = left + maxDepth(root.Left)
	}

	if root.Right != nil {
		right = right + maxDepth(root.Right)
	}

	if right > left {
		return right
	}
	return left
}
