package maximum_depth_of_binary_tree

import (
	"algorithm"
	"algorithm/binaryTree"
)

var res = -1<<63

func maxPathSum(root *binaryTree.TreeNode)int{
	dfs(root)
	return res
}

func dfs(root *binaryTree.TreeNode)int{
	if root == nil {
		return 0
	}
	leftMax := algorithm.Max(0,dfs(root.Left))
	rightMax := algorithm.Max(0,dfs(root.Right))

	res = algorithm.Max(res,root.Val+leftMax+rightMax)

	return root.Val+algorithm.Max(leftMax,rightMax)
}

