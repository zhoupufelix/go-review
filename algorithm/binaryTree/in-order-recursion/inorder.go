package in_order_recursion

import (
	"algorithm/binaryTree"
	"fmt"
)

func inOrder(root *binaryTree.TreeNode){
	if root == nil {
		return
	}
	inOrder(root.Left)
	fmt.Println(root.Val)
	inOrder(root.Right)
}
