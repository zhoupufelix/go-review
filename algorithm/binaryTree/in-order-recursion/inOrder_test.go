package in_order_recursion

import (
	"algorithm/binaryTree"
	"testing"
)

func TestInOrder(t *testing.T){
	root := binaryTree.CreateTreeNode()
	inOrder(root)
}
