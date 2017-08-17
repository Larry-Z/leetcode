package src

import (
	"container/list"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
递归深度优先
*/
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	dic := make(map[int]int)
	var treeDepthTravel func(root *TreeNode, curDepth int)
	treeDepthTravel = func(root *TreeNode, curDepth int) {
		if root == nil {
			return
		}
		if max, ok := dic[curDepth]; !ok || root.Val > max {
			dic[curDepth] = root.Val
		}
		fmt.Println("root.val=", root.Val)
		treeDepthTravel(root.Left, curDepth+1)
		treeDepthTravel(root.Right, curDepth+1)
	}
	treeDepthTravel(root, 0)
	ret := make([]int, len(dic))
	for dep, val := range dic {
		ret[dep] = val
	}
	return ret
}

/**
非递归，深度优先， 用map表示深度
*/
func largestValues2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	dic := make(map[int]int)

	stack := list.New()
	type Node struct {
		root  *TreeNode
		depth int
	}
	stack.PushFront(&Node{root, 0})
	for stack.Len() > 0 {
		ele := stack.Front()
		node := ele.Value.(*Node)
		stack.Remove(ele)
		if node.root == nil {
			continue
		}
		if max, ok := dic[node.depth]; !ok || node.root.Val > max {
			dic[node.depth] = node.root.Val
		}
		stack.PushFront(&Node{node.root.Left, node.depth + 1})
		stack.PushFront(&Node{node.root.Right, node.depth + 1})
	}
	ret := make([]int, len(dic))
	for dep, val := range dic {
		ret[dep] = val
	}
	return ret
}

/**
深度优先，只用数组
*/
func largestValues3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	ret := make([]int, 0, 100)
	stack := list.New()
	type Node struct {
		root  *TreeNode
		depth int
	}
	stack.PushFront(&Node{root, 0})
	depth := 0
	for stack.Len() > 0 {
		ele := stack.Front()
		node := ele.Value.(*Node)
		stack.Remove(ele)
		if node.root == nil {
			continue
		}
		if node.depth >= len(ret) {
			ret = append(ret, node.root.Val)
		} else if node.root.Val > ret[node.depth] {
			ret[node.depth] = node.root.Val
		}
		if node.depth > depth {
			depth = node.depth
		}
		stack.PushFront(&Node{node.root.Left, node.depth + 1})
		stack.PushFront(&Node{node.root.Right, node.depth + 1})
	}
	return ret
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Println(largestValues3(root))
}
