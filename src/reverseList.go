package src

import "fmt"

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	var prev *ListNode
	for p != nil {
		fmt.Println(*p)
		fmt.Println(prev)
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return prev
}

func main() {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 5}
	head.Next.Next = &ListNode{Val: 7}
	printLinkList(head)
	printLinkList(reverseList(head))
}
