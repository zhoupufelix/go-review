后序遍历：后序遍历左子树，后序遍历左子树，最后访问根节点

示例：
给定二叉树 [3,9,20,null,null,15,7]，

```go
    3
   / \
  9  20
    /  \
   15   7
```
打印结果 9 15 7 20 3

```go
func postOrder(root *TreeNode){ 
  if root == nil{
    return 
  }
  postOrder(root.Left)
  postOrder(root.Right)
  fmt.Println(root.Val)
}
```