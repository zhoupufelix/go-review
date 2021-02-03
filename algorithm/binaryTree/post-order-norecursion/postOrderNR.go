package post_order_norecursion

import (
	"algorithm/binaryTree"
	"fmt"
)

func postOrderNR(root *binaryTree.TreeNode){
	if root == nil {
		return
	}
	stack := make([]*binaryTree.TreeNode,0)
	var lastVisit *binaryTree.TreeNode

	for len(stack) != 0 || root != nil{
		for root != nil{
			stack = append(stack,root)
			root = root.Left
		}
		node := stack[len(stack)-1]

		//只有没有右节点 或者 右节点已经弹出 才可以弹出根节点
		if node.Right == nil ||  node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			fmt.Println(node.Val)
			lastVisit = node
		}else{
			root = node.Right
		}


	}
}