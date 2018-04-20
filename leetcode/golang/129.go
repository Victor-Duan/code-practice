/*
129. Sum Root to Leaf Numbers

Given a binary tree containing digits from 0-9 only, each root-to-leaf path could represent a number.

An example is the root-to-leaf path 1->2->3 which represents the number 123.

Find the total sum of all root-to-leaf numbers.

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
func sumNumbers(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return sumNumbersRoot(root, root.Val)
}

func sumNumbersRoot(root *TreeNode, val int) int {
    total := 0
    if root.Right == nil && root.Left == nil {
        return val
    }
    if root.Right != nil {
       total += sumNumbersRoot(root.Right, val * 10 + root.Right.Val) 
    }
    if root.Left != nil {
        total += sumNumbersRoot(root.Left, val * 10 + root.Left.Val)
    }
    fmt.Println(total)
    return total
}
