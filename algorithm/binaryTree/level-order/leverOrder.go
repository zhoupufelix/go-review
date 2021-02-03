package level_order

import (
	"algorithm/binaryTree"
	"fmt"
)

func levelOrder(root *binaryTree.TreeNode){
	if root == nil{
		return
	}
	queue := make([]*binaryTree.TreeNode,0)
	queue = append(queue,root)
	for len(queue) > 0{
		//pop
		node := queue[0]
		fmt.Println(node.Val)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue,node.Left)
		}
		if node.Right != nil {
			queue = append(queue,node.Right)
		}
	}
}
