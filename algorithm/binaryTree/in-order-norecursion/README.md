中序遍历：先中序遍历左子树，再遍历根节点，最后中序遍历右子树

示例：
给定二叉树 [3,9,20,null,null,15,7]，

```go
    3
   / \
  9  20
    /  \
   15   7
```
打印结果 9 3 15 20 7

```go
func inOrderNR(root *TreeNode){ 
  if root == nil{
    return 
  }
  stack = make([]*TreeNode,0)
  for root != nil || len(stack) > 0{
  	//一路向西
  	for root != nil{
        
    }
  }
}
```