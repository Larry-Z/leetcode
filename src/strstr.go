package main

import (
	"fmt"
	"strings"
)

func kmpNext(s string) []int {
	ret := make([]int, len(s))
	ret[0] = -1
	k, j := -1, 0
	for j < len(s)-1 {
		if k == -1 || s[k] == s[j] {
			k++
			j++
			ret[j] = k
		} else {
			k = ret[k]
		}
	}
	return ret
}

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(needle) == 1 {
		return strings.IndexByte(haystack, needle[0])
	}

	if len(haystack) == len(needle) {
		if haystack == needle {
			return 0
		}
		return -1
	}
	if len(needle) > len(haystack) {
		return -1
	}
	i, j := 0, 0
	next := kmpNext(needle)
	for i < len(haystack) && j < len(needle) {
		if j == -1 || haystack[i] == needle[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == len(needle) {
		return i - j
	}
	return -1
}

func main() {
	fmt.Println(strStr("abc", "bc"))
}
