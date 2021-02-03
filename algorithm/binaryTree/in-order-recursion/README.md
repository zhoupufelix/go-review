前序遍历：先访问左子树，再遍历中间，访问右子树

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
func inOrder(root *TreeNode){ 
  if root == nil{
    return 
  }
  inOrder(root.Left)
  fmt.Println(root.Val)
  inOrder(root.Right)
}
```