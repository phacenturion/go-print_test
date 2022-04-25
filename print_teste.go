package main

import (
	"fmt"
)

func Print_test() {
	fmt.Println("Hello world!")
}
func Str_test() string {
	return "another"
}

func main() {
	print(Str_test())
}
