/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

    dummy := &ListNode{Val : 0, Next : nil}
    tmp := dummy

    for (list1 != nil && list2 != nil) {

        if list1.Val < list2.Val {
           tmp.Next = list1
           list1 = list1.Next  
        } else {
            tmp.Next = list2
            list2 = list2.Next             
        } 
        tmp = tmp.Next
    }

    if list1 != nil {
       tmp.Next = list1 
    }   

    if list2 != nil {
        tmp.Next = list2
    } 

    return dummy.Next
}