package src

import "fmt"

func longestPalindrome(s string) int {
	m := [52]int{}
	for _, v := range s {
		k := v - 'A'
		if v > 'Z' {
			k = v - 6 - 'A'
		}
		m[k]++
	}
	ret := 0
	flag := false
	for _, v := range m {
		if v%2 == 0 {
			ret += v
		} else {
			ret += v - 1
			flag = true
		}
	}
	if flag {
		ret++
	}
	return ret
}

func main() {
	fmt.Println(longestPalindrome("aaaAaaaa") == 7)
}
