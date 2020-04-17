package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := map[string]int {
		"wjj": 39,
		"gxl": 28,
		"abc": 6,
	}

	for name, age := range ages {
		fmt.Println(name, age)
	}

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println(names)

	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name, ages[name])
	}

	fmt.Println(ages["xxx"], ages)
	
	if age, ok := ages["xxx"]; !ok {
		fmt.Println("name: xxx not exists!", age)
	}

	m := map[string]int{}
	// m := make(map[string]int)
	m["abc"] = 1
}