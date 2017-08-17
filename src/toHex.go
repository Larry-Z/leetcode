package src

import "fmt"

func toHex(num int) string {
	hex := "0123456789abcdef"
	ret := make([]byte, 8)
	if num == 0 {
		return "0"
	}
	count := 1
	for num != 0 && count <= 8 {
		ret[8-count] = hex[num&0xf]
		num = num >> 4
		count++
	}
	fmt.Println(ret, count)
	return string(ret[8-count+1:])
}

func main() {
	fmt.Println(toHex(1<<30 - 1))
}
