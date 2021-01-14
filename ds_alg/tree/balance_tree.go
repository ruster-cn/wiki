package tree

import (
	"log"
	"math"
)

/*
	## 什么是平衡二叉树
	左右子树的高度差不超过1,并且左右子树都是平衡二叉树

	思路：
		* 左子树不为空判断左子树是否为平衡二叉树，并计算max depth
		* 右子树不为空判断右子树是否为平衡二叉树，并计算max depth
		* 计算左右子树depth difference

	link:https://leetcode-cn.com/problems/balanced-binary-tree/
*/

func isBalanced(root *Node) bool {
	if root == nil {
		return true
	}
	leftDepth, rightDepth := 0, 0
	if root.Left != nil {
		if !isBalanced(root.Left) {
			return false
		}
		leftDepth = maxDepth(root.Left)
	}
	if root.Right != nil {
		if !isBalanced(root.Right) {
			return false
		}
		rightDepth = maxDepth(root.Right)
	}
	log.Printf("right depth is %d, left depth is %d", rightDepth, leftDepth)
	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		return false
	}
	return true
}
