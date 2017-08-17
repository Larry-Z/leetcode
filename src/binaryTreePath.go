package main

import (
	"fmt"
	"strconv"
)

func binaryTreePaths(root *TreeNode) []string {
	ret := []string{}
	var helper func(root *TreeNode, path string, ret *[]string)
	helper = func(root *TreeNode, path string, ret *[]string) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			path += strconv.Itoa(root.Val)
			*ret = append(*ret, path)
			return
		}
		path += fmt.Sprintf("%d->", root.Val)
		helper(root.Left, path, ret)
		helper(root.Right, path, ret)
	}
	helper(root, "", &ret)
	return ret
}

func main() {
	root := &TreeNode{Val: 1}
	root := &TreeNode{Val: 2}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 5}
	fmt.Println(binaryTreePaths(root))
}
