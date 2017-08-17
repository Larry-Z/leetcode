package main

import "fmt"

func findTarget(root *TreeNode, k int) bool {
	var findSingleNode func(cur, root *TreeNode, v int) bool
	findSingleNode = func(cur, root *TreeNode, v int) bool {
		if root == nil {
			return false
		}
		if root.Val == v && cur != root {
			return true
		}
		if root.Val > v {
			return findSingleNode(cur, root.Left, v)
		}
		return findSingleNode(cur, root.Right, v)
	}

	var findTargetHelper func(node *TreeNode, v int) bool
	findTargetHelper = func(node *TreeNode, v int) bool {
		if node == nil {
			return false
		}
		num := v - node.Val
		return findSingleNode(node, root, num) || findTargetHelper(node.Left, k) || findTargetHelper(node.Right, k)
	}
	return findTargetHelper(root, k)
}

func main() {
	list1 := "2,1,3"
	fmt.Println(findTarget(listToTree(list1), 4))
}
