package pre_order_recursion

import (
	"algorithm/binaryTree"
	"testing"
)

func TestPreOrderRecursion(t *testing.T){
	root := binaryTree.CreateTreeNode()
	preOrder(root)
}
