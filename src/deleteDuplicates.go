package main

func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    start, prev, count, head := new(ListNode), head, 1, head.Next
    p := start
    for head != nil {
        if head.Val == prev.Val {
            count++
        } else {
            if count == 1 {
                p.Next = prev
                p = p.Next
                prev.Next = nil
            }
            count = 1
        }
        prev = head
        head = head.Next
    }
    if count == 1 {
        p.Next = prev
        prev.Next = nil
    }
    return start.Next
}

func main() {
    head := sliceToList([]int{1,1})
    head = deleteDuplicates(head)
    printLinkList(head)
}
