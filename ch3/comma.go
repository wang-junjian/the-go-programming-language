package main

import "fmt"

func comma(s string) string {
	if len(s) <= 3 {
		return s
	}

	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func main() {
	fmt.Println(comma("123"))
	fmt.Println(comma("12345"))
	fmt.Println(comma("123456"))
	fmt.Println(comma("1234567"))
}