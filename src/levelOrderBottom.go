package main

import (
	"fmt"
)

func levelOrderBottom(root *TreeNode) [][]int {
	m := make(map[int][]int)

	depth := 0
	levelOrderBottom2(root, m, depth)
	ret := make([][]int, 0, len(m))
	for j := len(m) - 1; j >= 0; j-- {
		ret = append(ret, m[j])
	}
	return ret
}

func levelOrderBottom2(root *TreeNode, depthMap map[int][]int, depth int) {
	if root == nil {
		return
	}
	fmt.Println(depthMap, depth)
	if root.Left != nil {
		levelOrderBottom2(root.Left, depthMap, depth+1)
	}
	if root.Right != nil {
		levelOrderBottom2(root.Right, depthMap, depth+1)
	}
	var arr []int
	if v, ok := depthMap[depth]; ok {
		arr = v
	} else {
		arr = make([]int, 0)
	}
	arr = append(arr, root.Val)
	depthMap[depth] = arr
	fmt.Println(arr, depth)
}

func main() {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	fmt.Println(levelOrderBottom(root))

}
