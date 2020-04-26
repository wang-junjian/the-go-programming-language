package main

import (
	"fmt"
)

func mirroredQuery() {
	responses := make(chan string, 3)

	go func() {responses <- request("asia.gopl.io")}()
	go func() {responses <- request("europe.gopl.io")}()
	go func() {responses <- request("americas.gopl.io")}()

	return <-responses
}

func main() {
	response := mirroredQuery()
	fmt.Println(response)
}