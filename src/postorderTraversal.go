package main

import "fmt"

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := []int{}
	type Node struct {
		node *TreeNode
		top  bool
	}
	stack := []*Node{&Node{root, false}}
	for len(stack) > 0 {
		n := stack[len(stack)-1]
		if n.top == false {
			pop := true
			if n.node.Right != nil {
				stack = append(stack, &Node{n.node.Right, false})
				pop = false
			}
			if n.node.Left != nil {
				stack = append(stack, &Node{n.node.Left, false})
				pop = false
			}
			if pop {
				stack = stack[:len(stack)-1]
				ret = append(ret, n.node.Val)
			}
			n.top = true
		} else {
			ret = append(ret, n.node.Val)
			stack = stack[:len(stack)-1]
		}
	}
	return ret
}

func postOrder(node *TreeNode, ret *[]int) {
	if node == nil {
		return
	}
	postOrder(node.Left, ret)
	postOrder(node.Right, ret)
	*ret = append(*ret, node.Val)
}

func main() {
	root := listToTree("1,2,3,4,5,null,7,8")
	printTreePretty(root, 0)
	ret := []int{}
	postOrder(root, &ret)
	fmt.Println(ret)
	r := postorderTraversal(root)
	fmt.Println(r)
}
