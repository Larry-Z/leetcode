package main

import "fmt"

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	type Node struct {
		node *TreeNode
		sum  int
	}
	q := []*Node{&Node{root, 0}}
	sum := 0
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			s := q[i].sum*10 + q[i].node.Val
			if q[i].node.Left == nil && q[i].node.Right == nil {
				sum += s
				continue
			}
			if q[i].node.Left != nil {
				q = append(q, &Node{q[i].node.Left, s})
			}
			if q[i].node.Right != nil {
				q = append(q, &Node{q[i].node.Right, s})
			}
		}
		q = q[size:]
	}
	return sum
}

func main() {
	root := listToTree("1,2,3,9,null")
	printTreePretty(root, 0)
	fmt.Println(sumNumbers(root))
}
