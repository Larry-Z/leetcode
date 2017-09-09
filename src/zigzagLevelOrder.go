package main

import "fmt"

func zigzagLevelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    q := []*TreeNode{root}
    left := true
    ret := [][]int{}
    for len(q) > 0 {
        level := []int{}
        size := len(q)
        for i:=0; i<size; i++ {
            if left {
                level = append(level, q[i].Val)
            } else {
                level = append(level, q[size-1-i].Val)
            }
            if q[i].Left != nil {
                q = append(q, q[i].Left)
            }
            if q[i].Right != nil {
                q = append(q, q[i].Right)
            }
        }
        left = !left
        ret = append(ret, level)
        q = q[size:]
    }
    return ret
}

func main() {
    root := listToTree("3,9,20,null,null,15,7")
    fmt.Println(zigzagLevelOrder(root))
}
