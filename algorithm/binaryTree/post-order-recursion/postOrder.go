package post_order_recursion

import (
	"algorithm/binaryTree"
	"fmt"
)

func postOrder(root *binaryTree.TreeNode){
	if root == nil {
		return
	}
	postOrder(root.Left)
	postOrder(root.Right)
	fmt.Println(root.Val)
}
