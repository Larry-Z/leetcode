package src

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	mid := len(nums) / 2
	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

func main() {
	printTree(sortedArrayToBST([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
