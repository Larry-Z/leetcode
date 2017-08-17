package main

import "fmt"

func findLHS(nums []int) int {
	m := make(map[int]int, len(nums))
	ret := 0
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v1, ok := m[k+1]; ok {
			if v+v1 > ret {
				ret = v + v1
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(findLHS([]int{1, 3, 2, 2, 5, 2, 3, 7}))
}
