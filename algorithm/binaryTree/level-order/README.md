层序遍历：一层一层访问二叉树，利用队列，先进先出的数据结构。

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
func levelOrder(root *TreeNode){
	if root == nil{
		return
	}
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	for len(queue) > 0{
		//pop
		node := queue[0]
		fmt.Println(node.Val)
		queue = queue[1:]
		if node.Left != nil {
			queue = append(queue,node.Left)
		}
		if node.Right != nil {
			queue = append(queue,node.Right)
		}
	}
}
```