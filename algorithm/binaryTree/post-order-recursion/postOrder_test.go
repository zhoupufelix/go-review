package post_order_recursion

import (
	"algorithm/binaryTree"
	"testing"
)

func TestPostOrder(t *testing.T){
	root := binaryTree.CreateTreeNode()
	postOrder(root)
}
