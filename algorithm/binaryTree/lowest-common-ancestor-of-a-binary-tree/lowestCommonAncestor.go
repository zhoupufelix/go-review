package lowest_common_ancestor_of_a_binary_tree

import "algorithm/binaryTree"

func lowestCommonAncestor(root,p,q *binaryTree.TreeNode)*binaryTree.TreeNode{
	if root == nil  || root == p || root == q{
		return root
	}
	left  := lowestCommonAncestor(root.Left,p,q)
	right := lowestCommonAncestor(root.Right,p,q)

	if  left != nil && right != nil  {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	return nil
}