package main

import (
	"fmt"
	"math"
)

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}
	var first, second *TreeNode
	prev := &TreeNode{Val: math.MinInt32}

	var inOrder func(*TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		//fmt.Println(prev)
		if prev.Val >= node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		inOrder(node.Right)
	}
	inOrder(root)
	fmt.Println(first, second)
	first.Val, second.Val = second.Val, first.Val
}

func main() {
	root := listToTree("0,1,null,null")
	printTreePretty(root, 0)
	recoverTree(root)
	printTreePretty(root, 0)
}
