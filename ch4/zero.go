package main

import (
	"fmt"
)

func zero(ptr *[3]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}

func zero1(ptr *[3]byte) {
	*ptr = [3]byte{}
}

func main() {
	bytes := [3]byte{1,2,3}
	fmt.Printf("%d\n", bytes)
	zero1(&bytes)
	fmt.Printf("%d\n", bytes)
}