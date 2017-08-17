package src

import (
    "fmt"
)

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
    if root == nil {
        return nil
    }
    dic := make(map[string][]*TreeNode)
    var postOrder func(root *TreeNode) string
    postOrder = func(root *TreeNode) string {
        if root == nil {
            return ""
        }
        left := postOrder(root.Left)
        right := postOrder(root.Right)
        path := fmt.Sprintf("%s,%s,%d", left, right, root.Val)
        if list, ok := dic[path]; ok {
            dic[path] = append(list, root)
        } else {
            dic[path] = []*TreeNode{root}
        }
        return path
    }
    postOrder(root)
    //fmt.Println(dic)
    ret := make([]*TreeNode, 0)
    for _, list := range dic {
        if len(list) > 1 {
            ret = append(ret, list[0])
        }
    }
    return ret
}

func main() {
    root := listToTree("1,2,3,4,null,2,4,null,null,null,null,4,null,null,null,null,null")
    printTreePretty(root, 0)
    list := findDuplicateSubtrees(root)
    for _, node := range list {
        fmt.Printf("%d,", node.Val)
    }
}
