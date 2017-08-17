package main

import "fmt"

func isHappy(n int) bool {
	for n > 6 {
		ret := 0
		next := n
		for next > 0 {
			ret += (next % 10) * (next % 10)
			next = next / 10
		}
		n = ret
	}
	return n == 1
}

func main() {
	fmt.Println(isHappy(9999))
}
