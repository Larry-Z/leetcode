package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	count, bit := 1, 1
	for n > bit*9*count {
		n -= bit * 9 * count
		count *= 10
		bit++
		fmt.Println("n:", n, "count:", count, "bit:", bit)
	}
	div, mod := (n-1)/bit, (n-1)%bit
	return int((strconv.Itoa(count + div))[mod] - '0')
}

func main() {
	fmt.Println(findNthDigit(4))
}
