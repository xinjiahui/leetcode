package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	for cur.Next != nil {
		tmp1 := cur.Val
		tmp2 := cur.Next.Val
		//fmt.Println(tmp1, tmp2)
		if tmp1 == tmp2 {
			//fmt.Println("test333")
			cur.Next = cur.Next.Next

		} else {
			cur = cur.Next
		}

	}
	return head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 1,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	ret := reverseList(head)
	for ret != nil {
		fmt.Print(ret.Val, "->")
		ret = ret.Next
	}
}
