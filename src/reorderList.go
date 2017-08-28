package main

func reorderList(head *ListNode)  {
    if head == nil {
        return
    }
    // find the middle node
    p1, p2 := head, head
    for p2.Next != nil && p2.Next.Next != nil {
        p1 = p1.Next;
        p2 = p2.Next.Next;
    }

    // reverse second half
    var prev *ListNode
    for p1 != nil {
        tmp := p1.Next
        p1.Next = prev
        prev = p1
        p1 = tmp
    }

    // merge two list
    p0 := head
    for p0 != nil && prev != nil {
        n1, n2 := p0.Next, prev.Next
        p0.Next = prev
        prev.Next = n1
        p0 = n1
        prev = n2
    }
}

func main() {
    head := sliceToList([]int{1})
    printLinkList(head)
    reorderList(head)
    printLinkList(head)
}
