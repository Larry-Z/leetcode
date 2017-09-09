package main

import (
    "fmt"
)

func cal(num1, num2 int, op byte) int {
	switch op {
	case '+':
		return num1 + num2
	case '-':
		return num1 - num2
	case '*':
		return num1 * num2
	case '/':
		return num1 / num2
	}
	return 0
}

func parseNum(s string, start int) (int, int) {
	num, i := 0, start
	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		num = num*10 + int(s[i]-'0')
	}
	return num, i
}

func calculate(s string) int {
    if len(s) == 0 {
        return 0
    }
	opMap := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'(': -1,
	}
	numStack := []int{}
	opStack := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			num, j := parseNum(s, i)
			numStack = append(numStack, num)
			i = j - 1
			continue
		}
		if s[i] == '(' {
			opStack = append(opStack, s[i])
			continue
		}
		if s[i] == ')' {
			for opStack[len(opStack)-1] != '(' {
				opStack, numStack = calWithStack(opStack, numStack)
			}
			opStack = opStack[:len(opStack)-1]
			continue
		}
		op := s[i]
		for len(opStack) > 0 && opMap[opStack[len(opStack)-1]] >= opMap[op] {
			opStack, numStack = calWithStack(opStack, numStack)
		}
        opStack = append(opStack, op)
	}
	for len(opStack) > 0 {
		opStack, numStack = calWithStack(opStack, numStack)
	}
	return numStack[0]
}

func calWithStack(opStack []byte, numStack []int) ([]byte, []int) {
    fmt.Println(opStack, numStack)
    fmt.Println('(')
	op := opStack[len(opStack)-1]
	opStack = opStack[:len(opStack)-1]
	num2, num1 := numStack[len(numStack)-1], numStack[len(numStack)-2]
	numStack = numStack[:len(numStack)-2]
	num := cal(num1, num2, op)
	numStack = append(numStack, num)
	return opStack, numStack
}

func main() {
    expr := "2-1+2"
    fmt.Println(calculate(expr))
}
