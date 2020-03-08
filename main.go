package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var nums = []int{2, 5, 10, 1}
	fmt.Println(mergeTwoLists(nums, 27))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: 0, Next: nil}
	p := head
	p1, p2 := l1, l2
	for p1 != nil && p2 != nil {
		tmp := &ListNode{Val: 0, Next: nil}
		p.Next = tmp
		p = p.Next
		if p1.Val < p2.Val {
			tmp.Val = p1.Val
			p1 = p1.Next
		} else {
			tmp.Val = p2.Val
			p2 = p2.Next
		}
	}
	if p1 != nil {
		tmp := &ListNode{Val: p1.Val, Next: nil}
		p.Next, p1 = tmp, p1.Next
		p = p.Next
	}
	if p2 != nil {
		tmp := &ListNode{Val: p2.Val, Next: nil}
		p.Next, p2 = tmp, p2.Next
		p = p.Next
	}
	return head.Next
}
