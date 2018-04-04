/*
337. House Robber III

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// For every root, would it be better to steal from the current and the max of the sub subtree
// Or would it better to steal from the subtree
func rob(root *TreeNode) int {
    if root == nil {
        return 0
    } 
   
    // Leaf node, max value is the root value
    if root.Left == nil && root.Right == nil {
        return root.Val
    }
    
    // Otherwise, find two costs and compare
    costWithRoot := root.Val
    costWithoutRoot := 0
    if root.Right != nil {
        costWithRoot += rob(root.Right.Left) + rob(root.Right.Right)
        costWithoutRoot += rob(root.Right)
    }
    if root.Left != nil {
        costWithRoot += rob(root.Left.Right) + rob(root.Left.Left)
        costWithoutRoot += rob(root.Left)
    }
    
    if costWithRoot > costWithoutRoot {
        return costWithRoot
    }
    return costWithoutRoot
}
