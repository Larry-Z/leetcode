package main

import "fmt"

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{}
	stack := []*TreeNode{}
	n := root
	for len(stack) > 0 || n != nil {
		for n != nil {
			stack = append(stack, n)
			n = n.Left
		}
		n = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, n.Val)
		n = n.Right
	}
	return ret
}

func main() {
	root := listToTree("1,null,2,null,null,3,null")
	printTreePretty(root, 0)
	fmt.Println(inorderTraversal(root))
}
