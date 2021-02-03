package binaryTree



type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func CreateTreeNode()*TreeNode{
	leftNode   := &TreeNode{Val: 9 }
	rightLSub  := &TreeNode{Val: 15 }
	rightRSub  := &TreeNode{Val: 7 }
	rightNode  := &TreeNode{Val: 20 ,Left:rightLSub,Right: rightRSub}
	root       := &TreeNode{Val: 3,Left:leftNode,Right: rightNode}
	return root
}

