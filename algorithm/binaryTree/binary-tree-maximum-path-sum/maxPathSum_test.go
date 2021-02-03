package maximum_depth_of_binary_tree

import (
	"algorithm/binaryTree"
	"fmt"
	"testing"
)

func TestMaxPathSum(t *testing.T){
	root := binaryTree.CreateTreeNode()
	sum := maxPathSum(root)
	fmt.Println(sum)
}
