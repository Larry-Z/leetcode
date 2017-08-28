package main

func partition(head *ListNode, x int) *ListNode {
    if head == nil {
        return nil
    }
    var small, big, ret, bighead *ListNode
    for head != nil {
        if head.Val < x {
            if small == nil {
                small = head
                ret = head
            } else {
                small.Next = head
                small = head
            }
        } else {
            if big == nil {
                big = head
                bighead = head
            } else {
                big.Next = head
                big = big.Next
            }
        }
        head = head.Next
    }
    if ret != nil {
        small.Next = bighead
    } else {
        ret = bighead
    }
    if big != nil {
        big.Next = nil
    }
    return ret
}

func main() {
    head := sliceToList([]int{1})
    printLinkList(head)
    head = partition(head, 0)
    printLinkList(head)
}
