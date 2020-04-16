package main

import (
	"fmt"
	"strings"
)

func basename(s string) string {
	if n := strings.LastIndex(s, "/"); n >= 0 {
		s = s[n+1:]
	}

	if n := strings.LastIndex(s, "."); n>=0 {
		s = s[:n]
	}

	return s
}

func main() {
	fmt.Println(basename("a/b/c.go"))
	fmt.Println(basename("c.d.go"))
	fmt.Println(basename("abc"))
}