package maximum_depth_of_binary_tree

import (
	"algorithm"
	"algorithm/binaryTree"
)

func maxPath(root *binaryTree.TreeNode)int{
	if root == nil {
		return 0
	}
	return algorithm.Max(maxPath(root.Left),maxPath(root.Right))+1
}



