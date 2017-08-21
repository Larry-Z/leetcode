package main

import "fmt"

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	ret := []int{}
	for len(q) > 0 {
		size := len(q)
		ret = append(ret, q[len(q)-1].Val)
		for i := 0; i < size; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[size:]
	}
	return ret
}

func main() {
	str := "1,2,3,null,5,null,4"
	root := listToTree(str)
	printTreePretty(root, 0)
	fmt.Println(rightSideView(root))
}
