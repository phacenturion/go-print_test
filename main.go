package main

import (
	"fmt"
)

func print_test() {
	fmt.Println("Hello world!")
}
func str_test() string {
	return "another"
}
func main() {
	print(str_test())
}
