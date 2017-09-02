package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := []string{}
	dir := []byte{}
	for i := 0; i < len(path); {
		if path[i] == '/' {
			dir = dir[:0]
			i++
		}
		for i < len(path) && path[i] != '/' {
			dir = append(dir, path[i])
			i++
		}
		if len(dir) > 0 {
			s := string(dir)
			if s == "." {
				continue
			}
			if s == ".." {
				if len(stack) > 0 {
					//pop
					stack = stack[:len(stack)-1]
				}
				continue
			}
			stack = append(stack, s)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
