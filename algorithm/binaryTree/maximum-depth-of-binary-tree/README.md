给定一个二叉树,找出其最大深度
二叉树的深度为根结点到最远叶子节点的最长路径上的节点数。
说明: 叶子节点是指没有子节点的节点。

示例：
给定二叉树 [3,9,20,null,null,15,7]，

```go
    3
   / \
  9  20
    /  \
   15   7
```
返回它的最大深度 3 。
思路： 深度优先遍历

```go
func maxPath(root *TreeNode)int{
  if root == nil {
     return 0    
  }
  leftMax := maxPath(root.Left)
  rightmax := maxPath(root.Right)
  return Max(leftmax,rightmax)+1
}
 
```