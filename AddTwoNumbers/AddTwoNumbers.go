package main

import (
"fmt"
)

type ListNode struct {
	 Val int
     Next *ListNode
}	


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0
	dummyNode := &ListNode{Val: 0, Next: nil}
	cur := dummyNode

	for (l1 != nil) || (l2 != nil) || (carry != 0) {

		val1 := 0
		val2 := 0

		if l1 != nil {
			val1 , l1 = l1.Val , l1.Next            
		}
		if l2 != nil {
			val2, l2 = l2.Val, l2.Next           
		}

		sum := val1 + val2 + carry
		res := sum % 10
		carry = sum / 10

		newNode := &ListNode{Val: res, Next: nil}
		cur.Next = newNode
		cur = cur.Next    
        
	}
    fmt.Printf("loop \n")

	return dummyNode.Next

}

func main() {

	l1 := &ListNode{Val: 2, Next: nil}
	cur := l1
	n1a := &ListNode{Val: 4, Next: nil}
	cur.Next = n1a
	cur = cur.Next
	n1b := &ListNode{Val: 3, Next: nil}
	cur.Next = n1b
	cur = cur.Next
	
	l2 := &ListNode{Val: 5, Next: nil}
	cur = l2
	n2a := &ListNode{Val: 6, Next: nil}
	cur.Next = n2a
	cur = cur.Next
	n2b := &ListNode{Val: 4, Next: nil}
	cur.Next = n2b
	cur = cur.Next



	res := addTwoNumbers(l1, l2)

	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}

}	