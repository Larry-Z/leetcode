package src

func repeatedSubstringPattern(s string) bool {
	next := make([]int, len(s))
	for i := 2; i < len(s); i++ {

		if next[i-1] > 0 && s[i-1] == s[next[i]] {
			next[i] = next[i-1] + 1
		} else {
			next
		}
	}
}

func main() {
}
