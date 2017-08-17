package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := [26]int{}
	for k, v := range s {
		m[v-'a']++
		m[t[k]-'a']--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("aabb", "ccdd"))
}
