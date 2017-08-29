package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{Next: head}
	slow, fast := start, start
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return start.Next
}

func main() {
	head := sliceToList([]int{1, 2, 3, 4})
    printLinkList(head)
    head = removeNthFromEnd(head, 1)
    printLinkList(head)

}
