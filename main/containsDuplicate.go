package src

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	return len(m) == len(nums)
}

func main() {
}
