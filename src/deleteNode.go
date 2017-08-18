package main

import "fmt"

func delTreeNode (node *TreeNode) *TreeNode {
    if node == nil {
        return nil
    }
    if node.Left == nil {
        return node.Right
    }
    n := node.Left

    for n.Right != nil {
        n = n.Right
    }
    n.Right = node.Right
    return node.Left

}
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		return delTreeNode(root)
	}
	prev, n, isLeft := root, root.Left, true
	if key > root.Val {
		n, isLeft = root.Right, false
	}
	for n != nil {

		if key == n.Val {
			if isLeft {
				prev.Left = delTreeNode(n)
			} else {
				prev.Right = delTreeNode(n)
			}
			break
		}
		prev = n
		if key > n.Val {
			n = n.Right
			isLeft = false
		} else {
			n = n.Left
			isLeft = true
		}
	}
	return root
}

func main() {
	str := "17,12,22,9,15,19,23,3,10,14,16,18,20,null,null,null,null,null,null,13,null,null"
	root := listToTree(str)
	printTreePretty(root, 0)
	fmt.Println("-----------------------------")
	root = deleteNode(root, 17)
	fmt.Print(root)
	printTreePretty(root, 0)
}
