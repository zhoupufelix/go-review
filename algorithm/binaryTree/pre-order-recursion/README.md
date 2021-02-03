前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树

示例：
给定二叉树 [3,9,20,null,null,15,7]，

```go
    3
   / \
  9  20
    /  \
   15   7
```
打印结果 3 9 20 15 7

```go
func preOrder(root *TreeNode){ 
  if root == nil{
    return nil 
  }
  fmt.Println(root.Val)
  preOrder(root.Left)
  preOrder(root.Right)
}
```