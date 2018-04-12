/*
226. Invert Binary Tree

Invert a binary tree.

     4
   /   \
  2     7
 / \   / \
1   3 6   9
to

     4
   /   \
  7     2
 / \   / \
9   6 3   1
*/

// Invert the left and right subtrees
// Then set temp to left, left to right and right to temp

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    
    invR := invertTree(root.Right) 
    invL := invertTree(root.Left)

    root.Right = invL
    root.Left = invR
    
    return root
}
