package main

import "fmt"

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ret, sum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		fmt.Println("sum", sum, "ret", ret)
		sum = sum + nums[i]
		if ret < 0 && nums[i] > ret {
			ret = nums[i]
		}
		if sum < 0 {
			sum = 0
		} else {
			if sum > ret {
				ret = sum
			}
		}
	}
	return ret
}

func main() {
	fmt.Println(maxSubArray([]int{-1, 1, 2, 1}))
}
