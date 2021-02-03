package in_order_norecursion

import (
	"algorithm/binaryTree"
	"testing"
)

func TestInOrderNR(t *testing.T){
	root := binaryTree.CreateTreeNode()
	inOrderNR(root)
}
