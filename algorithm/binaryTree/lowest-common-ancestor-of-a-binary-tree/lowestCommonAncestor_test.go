package lowest_common_ancestor_of_a_binary_tree

import (
	"algorithm/binaryTree"
	"fmt"
	"testing"
)



func TestLowestCommonAncestor(t *testing.T){
	leftLSubL  := &binaryTree.TreeNode{Val: 7 }
	leftLSubR  := &binaryTree.TreeNode{Val: 4 }
	rightLSub  := &binaryTree.TreeNode{Val: 0 }
	rightRSub  := &binaryTree.TreeNode{Val: 8 }
	leftLSub   := &binaryTree.TreeNode{Val: 6 }
	leftRSub   := &binaryTree.TreeNode{Val: 2 ,Left:leftLSubL,Right: leftLSubR}
	leftNode   := &binaryTree.TreeNode{Val: 5 ,Left:leftLSub,Right: leftRSub}
	rightNode  := &binaryTree.TreeNode{Val: 1 ,Left:rightLSub,Right: rightRSub}
	root       := &binaryTree.TreeNode{Val: 3,Left:leftNode,Right: rightNode}
	ancestor := lowestCommonAncestor(root,leftNode,leftLSubR)
	fmt.Println(ancestor)
}
