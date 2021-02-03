package in_order_norecursion

import (
	"algorithm/binaryTree"
	"fmt"
)

func inOrderNR(root *binaryTree.TreeNode){
	stack := make([]*binaryTree.TreeNode,0)
	for len(stack) != 0 || root != nil{
		for root != nil{
			//push
			stack = append(stack,root)
			root = root.Left
		}
		//pop
		node := stack[len(stack)-1]
		fmt.Println(node.Val)
		stack = stack[:len(stack)-1]
		root = node.Right
	}
}