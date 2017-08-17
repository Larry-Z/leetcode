package main

import (
	"math"
)

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	findMaxIndex := func(arr []int) int {
		idx, max := -1, math.MinInt32
		for i, n := range nums {
			if n > max {
				idx, max = i, n
			}
		}
		return idx
	}
	idx := findMaxIndex(nums)
	root := &TreeNode{Val: nums[idx]}
	root.Left = constructMaximumBinaryTree(nums[0:idx])
	root.Right = constructMaximumBinaryTree(nums[idx+1:])
	return root
}

func main() {
	nums := []int{3, 2, 1, 6, 0, 5}
	printTree(constructMaximumBinaryTree(nums))
}
