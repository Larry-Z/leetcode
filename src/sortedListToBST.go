package main

func sortedListToBST(head *ListNode) *TreeNode {
	return listSliceToBST(head, nil)
}

func listSliceToBST(head, tail *ListNode) *TreeNode {
	if head == nil || head == tail {
		return nil
	}
	slow, fast := head, head
	for fast.Next != tail && fast.Next.Next != tail {
		slow, fast = slow.Next, fast.Next.Next
	}
	return &TreeNode{slow.Val, listSliceToBST(head, slow), listSliceToBST(slow.Next, tail)}
}

func main() {
    head := sliceToList([]int{1,2,3,4,5,6,7})
    printLinkList(head)
    root := sortedListToBST(head)
    printTreePretty(root, 0)
}
