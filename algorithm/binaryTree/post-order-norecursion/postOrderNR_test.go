package post_order_norecursion

import (
	"algorithm/binaryTree"
	"testing"
)

func TestPostOrderNR(t *testing.T){
	root := binaryTree.CreateTreeNode()
	postOrderNR(root)
}