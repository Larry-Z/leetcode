package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	len1, len2 := 0, 0
	for p1 != nil {
		len1++
		p1 = p1.Next
	}
	for p2 != nil {
		len2++
		p2 = p2.Next
	}
	var next *ListNode
	d := 0
	if len1 > len2 {
		p1, p2, d = l1, l2, len1-len2
	} else {
		p1, p2, d = l2, l1, len2-len1
	}
	for i := 0; i < d; i++ {
		node := &ListNode{Val: p1.Val, Next: next}
		next = node
		p1 = p1.Next
	}
	for p2 != nil {
		node := &ListNode{Val: p1.Val + p2.Val, Next: next}
		next, p1, p2 = node, p1.Next, p2.Next
	}
	carry := 0
	var prev *ListNode
	for next != nil {
		sum := next.Val + carry
		carry = sum / 10
		next.Val = sum % 10
		n := next.Next
		next.Next = prev
		prev = next
		next = n
	}
	if carry > 0 {
		return &ListNode{Val: carry, Next: prev}
	}
	return prev
}

func main() {
	l1 := sliceToList([]int{0})
	l2 := sliceToList([]int{0})
	printLinkList(l1)
	printLinkList(l2)
	printLinkList(addTwoNumbers(l1, l2))
}
