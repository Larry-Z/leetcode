package src

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	maxx, minn, end, begin := math.MinInt32, math.MaxInt32, -1, len(nums)
	for i := 0; i < len(nums); i++ {
		fmt.Printf("min:%d, max:%d, begin:%d, end:%d， num[i]:%d\n", minn, maxx, begin, end, nums[i])
		if nums[i] < maxx {
			end = i
		}
		maxx = max(maxx, nums[i])
		if nums[len(nums)-1-i] > minn {
			begin = len(nums) - i - 1
		}
		minn = min(minn, nums[len(nums)-1-i])
	}
	if end >= begin {
		return end - begin + 1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{1, 2, 3, 4, 5}))
}
