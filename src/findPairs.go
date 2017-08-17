package src

import "fmt"

func findPairs(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	dic := make(map[int][]int)
	for i, n := range nums {
		if arr, ok := dic[n]; ok {
			dic[n] = append(arr, i)
		} else {
			dic[n] = []int{i}
		}
	}
	fmt.Println("dic", dic)
	ret := make(map[string]bool)
	for i, n := range nums {
		aims := []int{n - k, n + k}
		for _, aim := range aims {
			if arr, ok := dic[aim]; ok {
				for _, j := range arr {
					if j > i {
						key := fmt.Sprintf("%d:%d", n, aim)
						if aim < n {
							key = fmt.Sprintf("%d:%d", aim, n)
						}
						ret[key] = true
						break
					}
				}
			}
		}
	}
	return len(ret)
}

func main() {
	fmt.Println(findPairs([]int{2, 8, 6, 9, 7, 4, 9, 0, 5, 4}, 1))
}
