package main

import "fmt"

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}
	start := &ListNode{Next: head}
	prev, pp, h, t, tail := start, head, head, head, head
	i := 1
	for head != nil {
		if i == m {
            pp = prev
			h = head
		}
		if i == n {
			t = head
			tail = head.Next
			break
		}
        prev = head
		head = head.Next
        i++
	}
    fmt.Println(pp, h, t, tail)
	p := h
	var pre *ListNode
	for p != tail {
		tmp := p.Next
		p.Next = pre
		pre = p
		p = tmp
	}
	pp.Next = t
	h.Next = tail
	return start.Next
}

func main() {
	head := sliceToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	head = reverseBetween(head, 1, 9)
	printLinkList(head)
}
