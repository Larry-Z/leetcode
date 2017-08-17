package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	l, h := 0, len(numbers)-1
	for l < h {
		sum := numbers[l] + numbers[h]
		if sum == target {
			return []int{l + 1, h + 1}
		}
		if numbers[l]+numbers[h] > target {
			h--
		} else {
			l++
		}
	}
	return []int{0, 0}
}

func main() {
	a := twoSum([]int{1, 4, 5, 9}, 6)
	fmt.Println(a)
}
