package src

import "fmt"

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricTree(root.Left, root.Right)
}

func isSymmetricTree(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l != nil && r != nil {
		return l.Val == r.Val && isSymmetricTree(l.Left, r.Right) && isSymmetricTree(l.Right, r.Left)
	}
	return false
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}
	root.Right.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 3}
	fmt.Println(isSymmetric(root))
}
