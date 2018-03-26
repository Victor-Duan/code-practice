/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func printLinkedList(head *ListNode) {
    for head != nil {
        fmt.Println(head.Val)
        head = head.Next
    } 
    fmt.Println("end")
}

func getLinkedListMiddleAndSplitList(head *ListNode) *ListNode{
    slow := head
    var slowPrev *ListNode = nil
    fast := head
    
    for fast != nil && fast.Next != nil {
        slowPrev = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    
    slowPrev.Next = nil;
    return slow
}

func sortedListToBST(head *ListNode) *TreeNode {
    // Base case: head is nil, return nil
    if head == nil {
        return nil
    }
    
    // Base case: List is 1 node long, just return head
    if head.Next == nil {
        return &TreeNode{
            Val: head.Val,
            Left: nil,
            Right: nil,
        }
    }
    
    // Start by finding the middle of the linked list; this is the root node
    middle := getLinkedListMiddleAndSplitList(head)
    
    rightList := middle.Next
    leftList := head
    middle.Next = nil
    
    printLinkedList(leftList)
    printLinkedList(rightList)
    
    // Then recursively call sortedListToBST on the left list and the right list
    leftRoot := sortedListToBST(leftList);
    rightRoot := sortedListToBST(rightList);
    
    // Then set Left and Right to the recursive function's return value
    return &TreeNode{
        Val: middle.Val,
        Left: leftRoot,
        Right: rightRoot,
    }
}
