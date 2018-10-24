/*
100. Same Tree
Given two binary trees, write a function to check if they are the same or not.

Two binary trees are considered the same if they are structurally identical and the nodes have the same value.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    } else if p == nil || q == nil && p != q {
        return false
    } else {
        if p.Val != q.Val {
            return false;
        } else {
            return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
        }
    }
}
