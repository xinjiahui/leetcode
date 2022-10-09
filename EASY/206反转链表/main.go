package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var per *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next //	把下一个节点缓存起来
		cur.Next = per
		per = cur
		cur = tmp

	}
	return per
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
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
