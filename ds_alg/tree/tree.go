package tree

//Node 定义树的叶子节点
type Node struct {
	Left  *Node
	Val   int
	Right *Node
}
