package src

import "fmt"

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}

	p := head
	for p.Next.Next != nil {
		p = p.Next
	}
	if head.Val != p.Next.Val {
		return false
	}
	p.Next = nil
	return isPalindrome(head.Next)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow = reverseList(slow)
	printLinkList(slow)
	printLinkList(head)
	for slow != nil {
		if slow.Val != head.Val {
			return false
		}
		slow = slow.Next
		head = head.Next
	}
	return true
}

func main() {
	head := &ListNode{Val: 10}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next = &ListNode{Val: 10}
	//fmt.Print(isPalindrome(head))
	//printLinkList(head)
	//printLinkList(reverseList(head))
	fmt.Println(isPalindrome2(head))
}
