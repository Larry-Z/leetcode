package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	start := &ListNode{Next: head}
	prev, p := start, head
	for p != nil {
		h := p
		for i := 0; i < k-1 && p != nil; i++ {
			p = p.Next
		}
		if p == nil {
			break
		}
		t, n := p, p.Next
		reverse(h, t)
		prev.Next = t
		prev = h
		h.Next = n
		p = n
	}
	return start.Next
}

func reverse(head, tail *ListNode) {
	var prev *ListNode
	for {
		tmp := head.Next
		head.Next = prev
		if head == tail {
			break
		}
		prev = head
		head = tmp
	}
}

func main() {
	head := sliceToList([]int{1, 2, 3, 4, 5, 6})
	head = reverseKGroup(head, 3)
	printLinkList(head)
}
