package src

import "fmt"

func findTilt1(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	l, s1 := findTilt1(root.Left)
	r, s2 := findTilt1(root.Right)
	s := abs(s1 - s2)

	return l + r + s, s1 + s2 + root.Val
}

func findTilt(root *TreeNode) int {
	l, _ := findTilt1(root)
	return l
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val:3}
	fmt.Println(findTilt(root))
}
