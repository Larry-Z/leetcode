package src

import "fmt"

func firstUniqChar(s string) int {
	m := [26]int{}
	for _, ch := range s {
		m[int(ch-'a')]++
	}
	for k, ch := range s {
		if m[int(ch-'a')] == 1 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("abacbdde"))
}
