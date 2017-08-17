package main

import "fmt"

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}
	i := 1
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if i == k {
			return root.Val
		}
		i++
		root = root.Right
	}
	return -1
}

func main() {
	list := "20,16,23,14,18,21,25,null,15,17,null"
	root := listToTree(list)
	printTreePretty(root, 0)
	fmt.Println(kthSmallest(root, 4))
}
