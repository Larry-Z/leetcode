package src

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	m := make(map[int]bool)
	for _, v := range nums1 {
		m[v] = true
	}
	ret := make([]int, 0)
	for _, v := range nums2 {
		if e, ok := m[v]; ok && e {
			ret = append(ret, v)
			m[v] = false
		}
	}
	return ret
}

func main() {
	a := []int{1, 3, 5, 7}
	b := []int{2, 4, 6, 8, 7, 7}
	fmt.Println(intersection(a, b))
}
