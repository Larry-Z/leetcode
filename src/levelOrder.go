package main

import "fmt"

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	ret := [][]int{}
	for len(q) > 0 {
		size := len(q)
		row := []int{}
		for i := 0; i < size; i++ {
			row = append(row, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		if len(row) > 0 {
			ret = append(ret, row)
		}
		q = q[size:]
	}
    return ret;
}

func main() {
    root := listToTree("3,9,20,null,null,15,7")
    printTreePretty(root, 0)
    fmt.Println(levelOrder(root))
}
