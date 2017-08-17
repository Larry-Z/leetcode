package main

import "fmt"

func majorityElement(nums []int) int {
	ret, count := 0, 0
	for _, v := range nums {
		if count == 0 {
			ret, count = v, 1
			continue
		}
		if v != ret {
			count--
		} else {
			count++
		}
	}
	return ret
}

func main() {
	fmt.Println(majorityElement([]int{1, 3, 5, 7, 3, 3, 3}))
}
