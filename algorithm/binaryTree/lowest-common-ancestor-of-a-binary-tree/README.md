##给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]


输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
输出: 3
解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

思路： 二叉树 后序遍历，如果左子树和右子树都不为空，那么根结点就是最近的公共祖先，另一种情况，如果左子树不为空，证明p,q都在左子树中。同理右子树。
。
```go
func lowestCommonAncestor(root,p,q *TreeNode)*TreeNode{
    if root == nil || root == p || root == q {
        return root
    }
    left := lowestCommonAncestor(root.Left,p,q)
    right := lowestCommonAncestor(root.Right,p,q)
    if left != nil && right != nil{
        return root
    }
    if left != nil {
       return left
    }
    if right != nil {
        return right
    }
    return nil
}
```