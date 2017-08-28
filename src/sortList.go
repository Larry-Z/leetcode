package main

import (
    "fmt"
    "math/rand"
)

func partition(head, tail *ListNode) *ListNode {
    if tail == nil {
        return nil
    }
    p, cur := head, head
    var prev *ListNode
    val := tail.Val
    for p != tail && p != nil {
        if p.Val < val {
            cur.Val, p.Val = p.Val, cur.Val
            prev = cur
            cur = cur.Next
        }
        p = p.Next
    }
    cur.Val, tail.Val = tail.Val, cur.Val
    return prev
}

func qsort(head, tail *ListNode) {
    if head == nil || tail == nil {
        return
    }
    if head == tail {
        return
    }
    if head.Next == tail && tail.Val < head.Val {
        head.Val, tail.Val = tail.Val, head.Val
        return
    }
    node := partition(head, tail)
    if node == nil {
        qsort(head.Next, tail)
        return
    }
    qsort(head, node)
    if node.Next == nil || node.Next.Next == nil {
        return
    }
    if node == tail || node.Next == tail || node.Next.Next == tail {
        return
    }
    qsort(node.Next.Next, tail)
}

func sortList(head *ListNode) *ListNode {
    tail := head
    for tail != nil {
        if tail.Next == nil {
            break
        }
        tail = tail.Next
    }
    qsort(head, tail)
    return head
}

func printList(head, tail *ListNode) {
    p := head
    for p != tail && p != nil {
        fmt.Printf("%d->", p.Val)
        p = p.Next
    }
    fmt.Printf("%d\n", tail.Val)
}

func main() {
    arr := []int{}
    for i := 0; i < 100; i++ {
        arr = append(arr, rand.Intn(200))
    }
    head := sliceToList(arr)
    printLinkList(head)
    sortList(head)
    printLinkList(head)
}
