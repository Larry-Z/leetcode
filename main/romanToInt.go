package src

import "fmt"

func romanToInt(s string) int {
	nums := make([]int, len(s))
	for k, ch := range s {
		switch ch {
		case 'I':
			nums[k] = 1
		case 'V':
			nums[k] = 5
		case 'X':
			nums[k] = 10
		case 'L':
			nums[k] = 50
		case 'C':
			nums[k] = 100
		case 'D':
			nums[k] = 500
		case 'M':
			nums[k] = 1000
		default:
			nums[k] = 0
		}
	}
	sum := 0
	for k, v := range nums {
		if k == len(nums)-1 {
			break
		}
		if v < nums[k+1] {
			sum -= v
		} else {
			sum += v
		}
		fmt.Println(sum)
	}
	sum += nums[len(nums)-1]
	return sum
}

func main() {
	fmt.Println(romanToInt("IV"))
}
