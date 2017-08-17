package src

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	l := len(num1)
	if len(num2) > len(num1) {
		l = len(num2)
	}
	ret := make([]byte, l+1)
	var carry byte
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; i-- {
		//fmt.Printf("i=%d,j=%d,l=%d\n", i, j, l)
		sum := carry
		if i >= 0 {
			sum += num1[i] - '0'
		}
		if j >= 0 {
			sum += num2[j] - '0'
		}
		if sum >= 10 {
			carry = 1
			sum -= 10
		} else {
			carry = 0
		}
		ret[l] = sum + '0'
		l--
		j--
	}
	if carry > 0 {
		ret[l] = 1 + '0'
		return string(ret)
	}
	return string(ret[1:])
}

func main() {
	s1 := "444"
	s2 := "333"
	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)
	fmt.Println(strconv.Itoa(i1 + i2))
	fmt.Printf(addStrings(s1, s2))
	fmt.Printf("%x", -1)
}
