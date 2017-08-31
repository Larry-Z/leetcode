package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	p := head
	l := 0
	for p != nil {
		p = p.Next
		l++
	}
	k = k % l
	if k == 0 {
		return head
	}
	fast, slow := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	var prevS, prevF *ListNode
	for fast != nil {
		prevF = fast
		fast = fast.Next
		prevS = slow
		slow = slow.Next
	}
	if prevS != nil {
		prevS.Next = nil
	}
	if prevF != nil {
		prevF.Next = head
	}
	return slow
}

func main() {
	head := sliceToList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	head = rotateRight(head, 3)
	printLinkList(head)
}
