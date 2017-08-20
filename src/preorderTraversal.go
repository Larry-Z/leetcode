package main

import (
	"fmt"
)

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	ret := make([]int, 0, 16)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, node.Val)
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return ret
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d,", root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

func main() {
	root := listToTree("1,2,3,4,5,6,7")
	printTreePretty(root, 0)
	preOrder(root)
	ret := preorderTraversal(root)
	fmt.Println(ret)
}
