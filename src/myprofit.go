package main

import "fmt"

func maxProfit(prices []int) int {
	ret := 0
	for k, v := range prices {
		if k+1 < len(prices) && v < prices[k+1] {
			ret += prices[k+1] - v
		}
	}
	return ret
}

func main() {
	fmt.Println(maxProfit([]int{1, 3, 5}))
}
