package main

import (
	"fmt"
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	strToComplex := func(str string) (int, int) {
		idx := strings.IndexByte(str, '+')
		real := str[0:idx]
		virt := str[idx+1 : len(str)-1]
		r, _ := strconv.Atoi(real)
		v, _ := strconv.Atoi(virt)
		return r, v
	}
	r1, v1 := strToComplex(a)
	r2, v2 := strToComplex(b)
	fmt.Println(r1, v1, r2, v2)
	r := r1*r2 - v1*v2
	v := r1*v2 + r2*v1
	return fmt.Sprintf("%d+%di", r, v)
}

func main() {
	fmt.Println(complexNumberMultiply("1+-1i", "0+0i"))
}
