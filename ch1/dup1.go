package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	// Mac Ctrl+D 退出
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Println()
	
	for line, n := range counts {
		if n>1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}