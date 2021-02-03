package maximum_depth_of_binary_tree

import (
	"algorithm/binaryTree"
	"testing"
)

func TestMaxPath(t *testing.T){
	root := binaryTree.CreateTreeNode()
	deep := maxPath(root)
	if deep != 3 {
		t.Error("expect 3")
	}
}

