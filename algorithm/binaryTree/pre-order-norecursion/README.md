前序遍历：先访问根节点，再前序遍历左子树，再前序遍历右子树
非递归：模拟一个栈，因为是先进后出的数据结构，所以先将右节点压入，再压入左节点。
所以先弹出左节点打印


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
func preOrderNR(root *TreeNode){ 
    if root == nil{
       return 
    }
    stack = make([]*TreeNode,0)
    stack = append(stack,root)
    for len(stack) != 0{
        //pop
        node := stack[len(stack)-1]
        fmt.Println(node.Val)
        stack = stack[:len(stack)-1]
        if node.Right != nil{
            stack = append(stack,node.Right)
        }        
        if node.Left != nil{
            stack = append(stack,node.Left)
        }   
    }
}

 
```