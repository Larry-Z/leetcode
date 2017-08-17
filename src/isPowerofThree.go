package main

import (
	"fmt"
	"strconv"
	"strings"
)

func isPowerOfThree(n int) bool {
	s := strconv.FormatInt(int64(n), 3)
	return s[0] == '1' && !strings.ContainsRune(s[1:], '2') && !strings.ContainsRune(s[1:], '1')
}

func isPowerOfTwo(n int) bool {
	s := 1 << 30
	return n > 0 && s%n == 0
}

func isPowerofTwo(n int) bool {
	return n > 0 && (n&(n-1) == 0)
}

func isPowerOfFour(num int) bool {
	n := 1
	for n <= num {
		if n == num {
			return true
		}
		n = n << 2
	}
	return false
}

func isPowerOfFour2(num int) bool {
	return num > 0 && (num&(num-1) == 0) && ((num & 0x55555555) != 0)
}

func main() {
	fmt.Println(isPowerOfFour2(32))
}
