/*
24. Swap Nodes in Pairs

Given a linked list, swap every two adjacent nodes and return its head.

For example,
Given 1->2->3->4, you should return the list as 2->1->4->3.

Your algorithm should use only constant space. You may not modify the values in the list, only nodes itself can be changed.
*/

// Keep two linked lists
// One linked list keeps track of the "odd" indexed elements
// Second keeps track of the "even" elements 
// Once this is constructed, then iterate through them and create the new list
func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
   
    oddList := head
    evenList := head.Next 
   
    curOddLoc := oddList
    curEvenLoc := evenList
    
    for head != nil {
        if head.Next != nil {
            head = head.Next.Next
        } else {
            break
        }
        
        curOddLoc.Next = head
        if head != nil && head.Next != nil {
            curEvenLoc.Next = head.Next
        } else {
            curEvenLoc.Next = nil
        }
        curOddLoc = curOddLoc.Next
        curEvenLoc = curEvenLoc.Next
    }
    
    curOddLoc = oddList
    curEvenLoc = evenList
    var evenTemp, oddTemp *ListNode

    for curEvenLoc != nil {
        // Track next
        evenTemp = curEvenLoc.Next
        oddTemp = curOddLoc.Next

        // swap values
        curEvenLoc.Next = curOddLoc
        if evenTemp != nil {
            curOddLoc.Next = evenTemp
        }

        // move to next nodes in the lists
        curEvenLoc = evenTemp
        curOddLoc = oddTemp
    }
    
    return evenList
}
