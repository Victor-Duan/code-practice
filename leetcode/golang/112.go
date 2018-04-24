/*
112. Path Sum

Given a binary tree and a sum, determine if the tree has a root-to-leaf path such that adding up all the values along the path equals the given sum.

Note: A leaf is a node with no children.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }
   
    if root.Left == nil && root.Right == nil {
        return root.Val == sum
    }
   
    hasSum := false
    if root.Left != nil {
        hasSum = hasSum || hasPathSum(root.Left, sum - root.Val) 
    }
    if root.Right != nil {
        hasSum = hasSum || hasPathSum(root.Right, sum - root.Val) 
    }
    return hasSum
}
