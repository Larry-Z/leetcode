package main

import (
	"fmt"
	"math"
)

/**
  Given a binary search tree with non-negative values,
  find the minimum absolute difference between values of any two nodes.
*/

func travel(root *TreeNode, min *int, prev **int) {
	if root == nil {
		return
	}
	travel(root.Left, min, prev)

	if *prev != nil && root.Val-**prev < *min {
		*min = root.Val - **prev
	}

	if *prev == nil {
		i := root.Val
		*prev = &i
	} else {
		**prev = root.Val
	}
	travel(root.Right, min, prev)
}

func getMinimumDifference(root *TreeNode) int {
	ret := math.MaxInt64
	var prev *int
	travel(root, &ret, &prev)
	return ret
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 3}
	fmt.Println(getMinimumDifference(root))
}
