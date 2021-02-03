package level_order

import (
	"algorithm/binaryTree"
	"testing"
)

func TestLevelOrder(t *testing.T){
	root := binaryTree.CreateTreeNode()
	levelOrder(root)
}
