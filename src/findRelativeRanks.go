package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	rank := make([]int, len(nums))
	copy(rank, nums)
	sort.Sort(sort.Reverse(sort.IntSlice(rank)))
	m := make(map[int]int, len(nums))
	for k, v := range rank {
		m[v] = k
	}
	ret := make([]string, len(nums))

	for k, v := range nums {
		r := m[v]
		var name string
		switch r {
		case 0:
			name = "Gold Medal"
		case 1:
			name = "Silver Medal"
		case 2:
			name = "Bronze Medal"
		default:
			name = strconv.Itoa(r + 1)
		}
		ret[k] = name
	}
	return ret
}

func main() {
	nums := []int{4, 7, 9, 8, 11}
	fmt.Println(findRelativeRanks(nums))
}
