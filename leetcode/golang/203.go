/*
203. Remove Linked List Elements

Remove all elements from a linked list of integers that have value val.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Keep track of prev, cur

func removeElements(head *ListNode, val int) *ListNode {
    // First make sure that head.Val is not equal to val
    for head != nil && head.Val == val{
        head = head.Next
    }
    if head == nil {
        return head
    }
   
    cur := head.Next
    prev := head
    for cur != nil {
        if cur.Val == val {
            prev.Next = cur.Next
        } else {
            prev = cur
        }
        cur = cur.Next
    }
    return head
}
