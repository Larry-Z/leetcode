package main

import (
    "fmt"
)

func isValidBST(root *TreeNode) bool {
    var prev *TreeNode
    ret := true
    var inOrder func(node *TreeNode)
    inOrder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inOrder(node.Left)
        if prev != nil && node.Val <= prev.Val {
            ret = false
        }
        prev = node
        inOrder(node.Right)
    }
    inOrder(root)
    return ret
}

func main() {
    root := listToTree("-2147483648")
    printTreePretty(root, 0)
    fmt.Println(isValidBST(root))
}
