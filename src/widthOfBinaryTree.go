package main

import "fmt"

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	type Node struct {
		node *TreeNode
		pos  int
	}
	q := []*Node{&Node{root, 0}}
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if q[i].node.Left != nil {
				q = append(q, &Node{q[i].node.Left, q[i].pos * 2})
			}
			if q[i].node.Right != nil {
				q = append(q, &Node{q[i].node.Right, q[i].pos*2 + 1})
			}
		}
		if q[size-1].pos-q[0].pos+1 > ret {
			ret = q[size-1].pos - q[0].pos + 1
		}
		q = q[size:]
	}
	return ret
}

func main() {
	root := listToTree("1,3,2,5,null,null,9,6,null,null,null,null,null,null,7")
	printTreePretty(root, 0)
	fmt.Println(widthOfBinaryTree(root))
}
