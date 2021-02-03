package pre_order_recursion

import (
	"algorithm/binaryTree"
	"fmt"
)

func preOrder(root *binaryTree.TreeNode){
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}
