package main

import "fmt"

func findBottomLeftValue(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var leftTravaseFunc func(root *TreeNode, depth int)
	curDepth, curVal := 0, root.Val

	leftTravaseFunc = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		fmt.Println("root.val", root.Val, "depth:", depth, "isLeft")
		if root.Left == nil && root.Right == nil && depth > curDepth {
			curDepth = depth
			curVal = root.Val
		}
		leftTravaseFunc(root.Left, depth+1)
		leftTravaseFunc(root.Right, depth+1)
	}

	leftTravaseFunc(root, 1)
	return curVal
}

func main() {
	root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	fmt.Println(findBottomLeftValue(root))
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 6}
	root.Right.Left.Left = &TreeNode{Val: 7}
	fmt.Println(findBottomLeftValue(root))
}
