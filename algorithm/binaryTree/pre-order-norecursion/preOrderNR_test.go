package pre_order_norecursion

import (
	"algorithm/binaryTree"
	"testing"
)

func TestPreOrderNR(t *testing.T){
	root := binaryTree.CreateTreeNode()
	preOrderNR(root)
}