package src

import "fmt"

func countBits(num int) []int {
	if num < 0 {
		return nil
	}
	ret := make([]int, 0, num+1)
	ret = append(ret, 0)
	for i := 1; i <= num; i++ {
		ret = append(ret, ret[i>>1]+(i&1))
	}
	return ret
}

func bitCount(n int) int {
	ret := 0
	for n > 0 {
		n = n & (n - 1)
		ret++
	}
	return ret
}

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("count(%d)=%d, ", i, bitCount(i))
		if i%4 == 0 {
			fmt.Println("")
		}
	}
	fmt.Println(countBits(100))
}
