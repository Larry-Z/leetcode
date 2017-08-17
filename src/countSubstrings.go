package src

import "fmt"

func countSubstrings(s string) int {
	if len(s) == 0 {
		return 0
	}
	count := 0
	for j := 0; j < len(s); j++ {
		if isPalindromic(s[:j+1]) {
			count++
		}
	}
	fmt.Println("s:", s, "j:", count)
	return count + countSubstrings(s[1:])
}

func isPalindromic(s string) bool {
	fmt.Println(s)
	low, high := 0, len(s)-1
	for low <= high {
		if s[low] != s[high] {
			return false
		}
		low++
		high--
	}
	return true
}

func main() {
	fmt.Println(isPalindromic("a"))
	fmt.Println(countSubstrings("aba"))
}
