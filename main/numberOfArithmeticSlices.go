package src

import "fmt"

func numberOfArithmeticSlices(A []int) int {
	if len(A) < 3 {
		return 0
	}
	count := 0
	if A[1]-A[0] == A[2]-A[1] {
		dec := A[1] - A[0]
		count++
		for i := 3; i < len(A); i++ {
			if A[i]-A[i-1] != dec {
				break
			}
			count++
		}
	}
	fmt.Println("header:", A[0], "count:", count)
	return count + numberOfArithmeticSlices(A[1:])
}

func numberOfArithmeticSlices2(A []int) int {
	if len(A) < 3 {
		return 0
	}
	ret := 0
	for j := 0; j < len(A)-2; j++ {
		count := 0
		if A[j+1]-A[j] == A[j+2]-A[j+1] {
			dec := A[j+1] - A[j]
			count++
			for i := j + 3; i < len(A); i++ {
				if A[i]-A[i-1] != dec {
					break
				}
				count++
			}
		}
		fmt.Println("header:", A[j], "count:", count)
		ret += count
	}
	//
	return ret
}

func numberOfArithmeticSlices3(A []int) int {
	if len(A) < 3 {
		return 0
	}
	for i := 0; i < len(A)-1; i++ {
		A[i] = A[i+1] - A[i]
	}
	A = A[:len(A)-1]
	l, ret := 1, 0
	fmt.Println(A)
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			ret += l * (l - 1) / 2
			l = 1
		} else {
			l++
		}
	}
	if l > 1 {
		ret += l * (l - 1) / 2
	}
	return ret
}

func main() {
	fmt.Println(numberOfArithmeticSlices3([]int{1, 2, 3, 4}))
}
