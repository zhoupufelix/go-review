后序遍历：先后序遍历左子树，再后序遍历右子树，最后访问根节点
非递归思路：对左节点一直压栈，只有当右节点为空或者右节点已经弹出，才能弹出当前节点

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
func postOrderNR(root *TreeNode){ 
  if root == nil{
    return 
  }
  stack = make([]*TreeNode,0)
  var lastVisit *TreeNode
  for len(stack) > 0 || root != nil{ 
    for root != nil{
       stack = appen(stack,root)
       root = root.Left
    }
    node := stack[len(stack)-1]
    if node.Right == nil || node.Right == lastVisit{ 
        //pop
        stack = stack[:len(stack)-1]
        fmt.Println(node.Val)
        lastVisit = node
    }else{
        root = root.Right
    }    
  }
}
```