package src

import "fmt"

func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	ret := make([]int, 0)
	fmt.Println(nums)
	for i, v := range nums {
		if v > 0 {
			ret = append(ret, i+1)
		}
	}
	return ret
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{1, 3, 2, 2, 5}))
}
