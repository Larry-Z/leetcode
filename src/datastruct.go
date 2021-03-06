package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func sliceToList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := &ListNode{Val: arr[0]}
	p := head
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{Val: arr[i]}
		p = p.Next
	}
	return head
}

func printLinkList(head *ListNode) {
	for head != nil {
		if head.Next == nil {
			fmt.Printf("%d", head.Val)
		} else {
			fmt.Printf("%d->", head.Val)
		}
		head = head.Next
	}
	fmt.Printf("\n")
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d, ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func printTreePretty(root *TreeNode, indent int) {
	if root == nil {
		return
	}
	printTreePretty(root.Right, indent+4)
	for i := 0; i < indent; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d\n", root.Val)
	printTreePretty(root.Left, indent+4)
}

func listToTree(str string) *TreeNode {
	list := strings.Split(str, ",")
	if len(list) == 0 {
		return nil
	}
	val, _ := strconv.Atoi(list[0])
	root := &TreeNode{Val: val}
	q, i := []*TreeNode{root}, 1
	for len(q) > 0 {
		size := len(q)
		for j := 0; j < size; j++ {
			if i >= len(list) || i+1 >= len(list) {
				break
			}
			if q[j] != nil {
				if list[i] != "null" {
					val, _ := strconv.Atoi(list[i])
					q[j].Left = &TreeNode{Val: val}
				}
				if list[i+1] != "null" {
					val, _ := strconv.Atoi(list[i+1])
					q[j].Right = &TreeNode{Val: val}
				}
				q = append(q, q[j].Left, q[j].Right)
			} else {
				q = append(q, nil, nil)
			}

			i += 2
		}
		q = q[size:]
	}
	return root
}

//func main() {
//	str := "1,2,3,4,5,6,7"
//    printTreePretty(listToTree(str), 0)
//
//}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
