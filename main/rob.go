package src

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func rob(nums []int) int {
	useLast, notUseLast := 0, 0
	for _, v := range nums {
		notUseLast, useLast = max(useLast, notUseLast), notUseLast+v
	}
	return max(useLast, notUseLast)
}

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
