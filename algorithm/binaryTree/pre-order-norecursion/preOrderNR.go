package pre_order_norecursion

import (
	"algorithm/binaryTree"
	"fmt"
)

func preOrderNR(root *binaryTree.TreeNode){
	if root == nil {
		return
	}
	stack := make([]*binaryTree.TreeNode,0)

	//push stack
	stack = append(stack,root)

	for len(stack) != 0{
		//pop
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(cur.Val)
		if cur.Right != nil {
			//push stack
			stack = append(stack,cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack,cur.Left)
		}
	}
	return
}
