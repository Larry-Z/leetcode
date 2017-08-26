package main

func insertionSortList(head *ListNode) *ListNode {
    p := head
    for p != nil {
        q := head
        for q != p {
            if q.Val > p.Val {
                p.Val, q.Val = q.Val, p.Val
            }
            q = q.Next
        }
        p = p.Next
    }
    return head
}

func selectSortList(head *ListNode) *ListNode  {
    p := head
    for p != nil {
        q := head
        for q != nil {
            if q.Val < p.Val {
                p.Val, q.Val = q.Val, p.Val
            }
            q = q.Next
        }
        p = p.Next
    }
    return head
}

func main() {
    head := sliceToList([]int{2, 7, 4, 5, 3, 6, 8, 1})
    printLinkList(head)
    head = insertionSortList(head)
    printLinkList(head)
}
