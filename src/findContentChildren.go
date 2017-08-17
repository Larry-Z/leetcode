package src

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	i := 0
	for _, v := range s {
		if i == len(g) {
			return ret
		}
		if v >= g[i] {
			ret++
			i++
		}
	}
	return ret
}

func main() {
	a := []int{1, 3, 4, 5}
	b := []int{1, 2, 4, 5}
	fmt.Println(findContentChildren(a, b))
}
