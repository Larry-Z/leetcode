package src

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	m := [26]int{}
	for _, v := range magazine {
		m[v-'a']++
	}
	for _, v := range ransomNote {
		if m[v-'a'] == 0 {
			return false
		}
		m[v-'a']--
	}
	return true
}

func main() {
	fmt.Println(canConstruct("aac", "aab"))
}
